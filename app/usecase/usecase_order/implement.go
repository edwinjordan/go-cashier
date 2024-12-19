package usecase_order

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
)

type UseCaseImpl struct {
	CustomerOrderRepository       repository.CustomerOrderRepository
	CustomerOrderDetailRepository repository.CustomerOrderDetailRepository
	BoothStockRepository          repository.BoothStockRepository
	VarianRepository              repository.VarianRepository
	TempCartRepository            repository.TempCartRepository
	TransactionRepository         repository.TransactionRepository
	TransactionDetailRepository   repository.TransactionDetailRepository
	UserRepository                repository.UserRepository
	Validate                      *validator.Validate
}

func NewUseCase(
	orderRepo repository.CustomerOrderRepository,
	orderDetailRepo repository.CustomerOrderDetailRepository,
	varianRepo repository.VarianRepository,
	tempCartRepo repository.TempCartRepository,
	transRepo repository.TransactionRepository,
	transDetailRepo repository.TransactionDetailRepository,
	stockBoothRepo repository.BoothStockRepository,
	userRepo repository.UserRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		Validate:                      validate,
		CustomerOrderRepository:       orderRepo,
		CustomerOrderDetailRepository: orderDetailRepo,
		VarianRepository:              varianRepo,
		TempCartRepository:            tempCartRepo,
		TransactionRepository:         transRepo,
		TransactionDetailRepository:   transDetailRepo,
		BoothStockRepository:          stockBoothRepo,
		UserRepository:                userRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	dataParent := dataRequest["parent"].(map[string]interface{})

	/* check apakah masih ada sisa untuk product tsb */
	for _, v := range dataRequest["detail"].([]interface{}) {
		dt := v.(map[string]interface{})
		detail, _ := controller.VarianRepository.FindById(r.Context(), dt["product_varian_id"].(string))
		/* jika ada salah satu barang memiliki sisa kurang dari yang dipesan maka langsung batalkan pesanan */
		if detail.ProductVarianQtyLeft < int(dt["product_qty"].(float64)) {
			panic(exceptions.NewBadRequestError("Tidak dapat membuat order karena salah satu barang dikeranjang tidak memiliki jumlah tersisa yang cukup"))
		}
	}

	/* masukkan ke order */
	dataResponse := controller.CustomerOrderRepository.Create(r.Context(), entity.CustomerOrder{
		OrderInvNumber:    controller.CustomerOrderRepository.GenInvoice(r.Context()),
		OrderCustomerId:   dataParent["order_customer_id"].(string),
		OrderAddressId:    dataParent["order_address_id"].(string),
		OrderDeliveryType: dataParent["order_delivery_type"].(string),
		OrderTotalItem:    int(dataParent["order_total_item"].(float64)),
		OrderSubtotal:     dataParent["order_subtotal"].(float64),
		OrderDiscount:     dataParent["order_discount"].(float64),
		OrderTotal:        dataParent["order_total"].(float64),
		OrderNotes:        html.EscapeString(dataParent["order_notes"].(string)),
	})

	/* masukkan ke table detail */
	for _, v := range dataRequest["detail"].([]interface{}) {
		dt := v.(map[string]interface{})
		controller.CustomerOrderDetailRepository.Create(r.Context(), entity.CustomerOrderDetail{
			OrderDetailParentId:        dataResponse.OrderId,
			OrderDetailProductVarianId: dt["product_varian_id"].(string),
			OrderDetailQty:             int(dt["product_qty"].(float64)),
			OrderDetailPrice:           dt["product_price"].(float64),
			OrderDetailSubtotal:        (dt["product_qty"].(float64) * dt["product_price"].(float64)),
		})

		/* masukkan ke temporary  */
		controller.TempCartRepository.Create(r.Context(), entity.TempCart{
			TempCartProductVarianId: dt["product_varian_id"].(string),
			TempCartUserId:          dataParent["order_customer_id"].(string),
			TempCartQty:             int(dt["product_qty"].(float64)),
			TempCartOrderId:         dataResponse.OrderId,
		})
	}

	/* get ada admin */
	users := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
		UserHasMobileAccess: 1,
	})
	for _, v := range users {
		dt := helpers.GetFCMToken(v.UserId)
		if len(dt) > 0 {
			helpers.SendFCM(dt, map[string]interface{}{
				"title": "Pesanan Baru",
				"body":  "Ada pesanan baru senilai Rp. " + fmt.Sprint(int(dataResponse.OrderTotal)) + ", klik untuk melihat detail",
				"data": map[string]interface{}{
					"id":   dataResponse.OrderId,
					"type": "order",
				},
			})
		}
	}
	/* kirim notif ke admin */

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataRequest := entity.CustomerOrder{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	_, err = controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataRequest.OrderId = id
	dataResponse := controller.CustomerOrderRepository.Update(r.Context(), dataRequest, "*", entity.CustomerOrder{
		OrderId: id,
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	_, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.CustomerOrderRepository.Delete(r.Context(), entity.CustomerOrder{
		OrderId: id,
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataResponse, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	where := entity.CustomerOrder{}

	if vars.Get("order_customer_id") != "" {
		where.OrderCustomerId = vars.Get("order_customer_id")
	}
	if vars.Get("order_status") != "" {
		status, _ := strconv.Atoi(vars.Get("order_status"))
		where.OrderStatus = status
	}

	Qlimit := vars.Get("limit")
	Qoffset := vars.Get("offset")

	if Qlimit == "" {
		Qlimit = "10"
	}

	if Qoffset == "" {
		Qoffset = "0"
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset, _ := strconv.Atoi(Qoffset)

	nextOffset := limit + offset

	conf := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")

	dataResponse := controller.CustomerOrderRepository.FindSpesificData(r.Context(), where, conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	where := entity.ViewOrderDetail{}

	if vars.Get("order_detail_id") != "" {
		where.OrderDetailId = vars.Get("order_detail_id")
	}

	if vars.Get("order_detail_parent_id") != "" {
		where.OrderDetailParentId = vars.Get("order_detail_parent_id")
	}

	dataResponse := controller.CustomerOrderDetailRepository.FindSpesificData(r.Context(), where)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) OrderCanceled(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataRequest := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &dataRequest)
	order, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	/* ambil data dari jwt jika bukan kasir maka hanya bisa membatalkan jika statusnya belum diproses */
	data := r.Context().Value("userslogin").(jwt.MapClaims)
	/* check apakah orderan sudah di proses */
	if order.OrderStatus == 2 && data["HasAccessCashier"].(float64) == 0 {
		panic(exceptions.NewBadRequestError("Tidak dapat membatalkan pesanan karena sudah diproses, silahkan hubungi kasir untuk melakukan pembatalan"))
	}
	notes := ""
	cancelBy := order.OrderCustomerId
	if data["HasAccessCashier"].(float64) == 0 {
		notes = "Dibatalkan oleh pelanggan."
	} else {
		notes = dataRequest["message"].(string)
		cancelBy = data["UserId"].(string)
	}

	controller.CustomerOrderRepository.Update(r.Context(), entity.CustomerOrder{
		OrderStatus: 3,
		OrderFinishedDatetime: sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		},
		OrderFinishedBy:  cancelBy,
		OrderCancelNotes: notes,
	}, []string{"order_status", "order_finished_datetime", "order_cancel_notes", "order_finished_by"}, entity.CustomerOrder{
		OrderId: id,
	})

	/* hapus data temporary cart */
	controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
		TempCartOrderId: id,
	})

	/* kirim notif ke kasir dan pelanggan */
	users := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
		UserHasMobileAccess: 1,
	})
	for _, v := range users {
		dt := helpers.GetFCMToken(v.UserId)
		if len(dt) > 0 {
			helpers.SendFCM(dt, map[string]interface{}{
				"title": "Pesanan Dibatalkan",
				"body":  "Pesanan dengan nomor order " + order.OrderInvNumber + " telah dibatalkan, klik untuk melihat detail",
				"data": map[string]interface{}{
					"id":   order.OrderId,
					"type": "order",
				},
			})
		}
	}

	dtCust := helpers.GetFCMToken(order.OrderCustomerId)
	if len(dtCust) > 0 {
		helpers.SendFCM(dtCust, map[string]interface{}{
			"title": "Pesanan Dibatalkan",
			"body":  "Pesanan dengan nomor order " + order.OrderInvNumber + " telah dibatalkan, klik untuk melihat detail",
			"data": map[string]interface{}{
				"id":   order.OrderId,
				"type": "order",
			},
		})
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil membatalkan pesanan",
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) OrderProcessed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	reqData := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &reqData)
	// helpers.ReadFromRequestBody(r, &dataRequest)
	// err := controller.Validate.Struct(dataRequest)
	// helpers.PanicIfError(err)
	dataRequest := entity.CustomerOrder{
		OrderStatus: int(reqData["order_status"].(float64)),
	}
	order, err := controller.CustomerOrderRepository.FindById(r.Context(), id)

	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataUser := r.Context().Value("userslogin").(jwt.MapClaims)
	selectField := []string{"order_status", "order_processed_datetime", "order_processed_by"}
	message := ""
	fcmMessage := "Pesanan dengan nomor order " + order.OrderInvNumber + " telah diproses, klik untuk melihat detail"
	fcmTitle := "Pesanan Diproses"
	if dataRequest.OrderStatus == 2 {
		dataRequest.OrderProcessedDatetime = sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		}
		dataRequest.OrderProcessedBy = dataUser["UserId"].(string)
		message = "Berhasil memproses pesanan"
	} else if dataRequest.OrderStatus == 1 {
		fcmMessage = "Pesanan dengan nomor order " + order.OrderInvNumber + " telah selesai, klik untuk melihat detail"
		fcmTitle = "Pesanan Selesai"

		selectField = []string{"order_status", "order_finished_datetime", "order_finished_by"}
		dataRequest.OrderFinishedDatetime = sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		}
		dataRequest.OrderFinishedBy = dataUser["UserId"].(string)
		message = "Berhasil menyelesaikan pesanan"

		/* masukkan ke table transaksi */
		parent := reqData["parent"].(map[string]interface{})
		qtyTotal := 0
		for _, v := range reqData["cart_detail"].([]interface{}) {
			dt := v.(map[string]interface{})
			qtyTotal += int(dt["product_qty"].(float64))
		}

		invoice := controller.TransactionRepository.GenInvoice(r.Context())

		/* insert into parent data */
		dataTransaction := entity.Transaction{
			TransUserId:        parent["user_id"].(string),
			TransInvoice:       invoice,
			TransOrderId:       id,
			TransQtyTotal:      qtyTotal,
			TransProductTotal:  len(reqData["cart_detail"].([]interface{})),
			TransSubtotal:      parent["total_price"].(float64),
			TransDiscount:      parent["total_discount"].(float64),
			TransTotal:         parent["total_price"].(float64) - parent["total_discount"].(float64),
			TransReceivedTotal: parent["total_receive"].(float64),
			TransRefundTotal:   parent["total_receive"].(float64) - (parent["total_price"].(float64) - parent["total_discount"].(float64)),
			TransCustomerId:    parent["customer_id"].(string),
			TransStatus:        1,
		}
		trans := controller.TransactionRepository.Create(r.Context(), dataTransaction)

		/* input detail */
		for _, v := range reqData["cart_detail"].([]interface{}) {

			dt := v.(map[string]interface{})
			subtotal := dt["product_qty"].(float64) * dt["product_price"].(float64)
			transDetail := entity.TransactionDetail{
				TransDetailParentId:        trans.TransId,
				TransDetailProductVarianId: dt["product_varian_id"].(string),
				TransDetailQty:             int(dt["product_qty"].(float64)),
				TransDetailPrice:           dt["product_price"].(float64),
				TransDetailSubtotal:        subtotal,
			}
			controller.TransactionDetailRepository.Create(r.Context(), transDetail)

			/* input kartu stok */
			varian, _ := controller.VarianRepository.FindById(r.Context(), dt["product_varian_id"].(string))
			lastStock := varian.ProductVarianQtyBooth - int(dt["product_qty"].(float64))
			controller.BoothStockRepository.Create(r.Context(), entity.StockBooth{
				ProductStokProductVarianId: dt["product_varian_id"].(string),
				ProductStokFirstQty:        varian.ProductVarianQtyBooth,
				ProductStokQty:             int(dt["product_qty"].(float64)),
				ProductStokLastQty:         lastStock,
				ProductStokJenis:           "keluar",
				ProductStokPegawaiId:       parent["user_id"].(string),
			})

			/* kurangi stok booth */
			controller.VarianRepository.UpdateStock(r.Context(), dt["product_varian_id"].(string), lastStock)
		}

		/* hapus table temp cart */
		controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
			TempCartOrderId: id,
		})
		/* sampai sini */
	} else {
		panic(exceptions.NewBadRequestError("Kesalahan tidak diketahui, silahkan hubungi administrator anda"))
	}
	controller.CustomerOrderRepository.Update(r.Context(), dataRequest, selectField, entity.CustomerOrder{
		OrderId: id,
	})

	/* kirim notif ke pelanggan */

	dtCust := helpers.GetFCMToken(order.OrderCustomerId)
	if len(dtCust) > 0 {
		helpers.SendFCM(dtCust, map[string]interface{}{
			"title": fcmTitle,
			"body":  fcmMessage,
			"data": map[string]interface{}{
				"id":   order.OrderId,
				"type": "order",
			},
		})
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: message,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
