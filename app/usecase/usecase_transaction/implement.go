package usecase_transaction

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
)

type UseCaseImpl struct {
	TransactionRepository         repository.TransactionRepository
	TransactionDetailRepository   repository.TransactionDetailRepository
	TempCartRepository            repository.TempCartRepository
	BoothStockRepository          repository.BoothStockRepository
	VarianRepository              repository.VarianRepository
	CustomerOrderRepository       repository.CustomerOrderRepository
	CustomerOrderDetailRepository repository.CustomerOrderDetailRepository
	Validate                      *validator.Validate
}

func NewUseCase(
	transRepo repository.TransactionRepository,
	transDetailRepo repository.TransactionDetailRepository,
	tempCartRepo repository.TempCartRepository,
	stockBoothRepo repository.BoothStockRepository,
	varianRepo repository.VarianRepository,
	orderRepo repository.CustomerOrderRepository,
	orderDetailRepo repository.CustomerOrderDetailRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		Validate:                      validate,
		TransactionRepository:         transRepo,
		TransactionDetailRepository:   transDetailRepo,
		TempCartRepository:            tempCartRepo,
		BoothStockRepository:          stockBoothRepo,
		CustomerOrderRepository:       orderRepo,
		CustomerOrderDetailRepository: orderDetailRepo,
		VarianRepository:              varianRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	parent := dataRequest["parent"].(map[string]interface{})
	qtyTotal := 0
	for _, v := range dataRequest["cart_detail"].([]interface{}) {
		dt := v.(map[string]interface{})
		qtyTotal += int(dt["product_qty"].(float64))
	}

	invoice := controller.TransactionRepository.GenInvoice(r.Context())

	orderId := ""
	if parent["customer_id"].(string) != "" {

		dataOrder := controller.CustomerOrderRepository.Create(r.Context(), entity.CustomerOrder{
			OrderInvNumber:    controller.CustomerOrderRepository.GenInvoice(r.Context()),
			OrderCustomerId:   parent["customer_id"].(string),
			OrderAddressId:    "",
			OrderDeliveryType: "AMBIL",
			OrderTotalItem:    len(dataRequest["cart_detail"].([]interface{})),
			OrderSubtotal:     parent["total_price"].(float64),
			OrderDiscount:     parent["total_discount"].(float64),
			OrderTotal:        parent["total_price"].(float64) - parent["total_discount"].(float64),
			OrderNotes:        "Generate automatically by system",
		})

		controller.CustomerOrderRepository.Update(r.Context(), entity.CustomerOrder{
			OrderStatus: 1,
		}, "order_status", entity.CustomerOrder{OrderId: dataOrder.OrderId})

		orderId = dataOrder.OrderId
		/* masukkan ke table detail */
		for _, v := range dataRequest["cart_detail"].([]interface{}) {
			dt := v.(map[string]interface{})
			controller.CustomerOrderDetailRepository.Create(r.Context(), entity.CustomerOrderDetail{
				OrderDetailParentId:        dataOrder.OrderId,
				OrderDetailProductVarianId: dt["product_varian_id"].(string),
				OrderDetailQty:             int(dt["product_qty"].(float64)),
				OrderDetailPrice:           dt["product_price"].(float64),
				OrderDetailSubtotal:        dt["product_qty"].(float64) * dt["product_price"].(float64),
			})
		}
	}

	/* insert into parent data */
	dataTransaction := entity.Transaction{
		TransUserId:        parent["user_id"].(string),
		TransInvoice:       invoice,
		TransOrderId:       orderId,
		TransQtyTotal:      qtyTotal,
		TransProductTotal:  len(dataRequest["cart_detail"].([]interface{})),
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
	for _, v := range dataRequest["cart_detail"].([]interface{}) {

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
		TempCartUserId: parent["user_id"].(string),
	})

	/* ambil data transaksi */
	dataResponse, _ := controller.TransactionRepository.FindById(r.Context(), trans.TransId)

	/* kirim notif ke pelanggan */

	if parent["customer_id"].(string) != "" {

		dtCust := helpers.GetFCMToken(parent["customer_id"].(string))
		if len(dtCust) > 0 {
			helpers.SendFCM(dtCust, map[string]interface{}{
				"title": "Pesanan Selesai",
				"body":  "Pesanan dengan nomor transaksi " + dataResponse.TransInvoice + " telah selesai, klik untuk melihat detail",
				"data": map[string]interface{}{
					"id":   orderId,
					"type": "order",
				},
			})
		}
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["TransactionId"]
	dataRequest := entity.Transaction{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	Transaction, err := controller.TransactionRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataRequest.TransId = id
	dataRequest.TransCreateAt = Transaction.TransCreateAt
	dataResponse := controller.TransactionRepository.Update(r.Context(), dataRequest, id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["TransactionId"]
	dataResponse, err := controller.TransactionRepository.FindById(r.Context(), id)
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
	today := time.Now()
	where := entity.Transaction{}

	if vars.Get("trans_user_id") != "" {
		where.TransUserId = vars.Get("trans_user_id")
	}

	if vars.Get("trans_customer_id") != "" {
		where.TransCustomerId = vars.Get("trans_customer_id")
	}

	date := vars.Get("date")
	dataDate := map[string]interface{}{
		"startDate": "",
		"endDate":   "",
	}

	if date != "" {
		switch date {
		case "today":
			dataDate["startDate"] = today.Format("2006-01-02")
			// dataDate["endDate"] = today.Format("2006-01-02")
		case "yesterday":
			dataDate["startDate"] = today.AddDate(0, 0, -1).Format("2006-01-02")
			// dataDate["endDate"] = today.AddDate(0, 0, -1).Format("2006-01-02")
		case "7":
			dataDate["startDate"] = today.AddDate(0, 0, -7).Format("2006-01-02")
			dataDate["endDate"] = today.Format("2006-01-02")
		case "this_month":
			dataDate["startDate"] = today.Format("2006-01")
		case "30":
			dataDate["startDate"] = today.AddDate(0, 0, -30).Format("2006-01-02")
			dataDate["endDate"] = today.Format("2006-01-02")
		}
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
		"limit":      limit,
		"offset":     offset,
		"date":       dataDate,
		"typeofDate": date,
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")

	dataResponse := controller.TransactionRepository.FindSpesificData(r.Context(), where, conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetTransDetail(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	where := entity.ViewTransactionDetail{}

	if vars.Get("trans_detail_id") != "" {
		where.TransDetailId = vars.Get("trans_detail_id")
	}

	if vars.Get("trans_detail_parent_id") != "" {
		where.TransDetailParentId = vars.Get("trans_detail_parent_id")
	}

	dataResponse := controller.TransactionDetailRepository.FindSpesificData(r.Context(), where)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetTransactionSummary(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	today := time.Now()

	date := vars.Get("date")
	dataDate := map[string]interface{}{
		"startDate": "",
		"endDate":   "",
	}

	if date != "" {
		switch date {
		case "today":
			dataDate["startDate"] = today.Format("2006-01-02")
			// dataDate["endDate"] = today.Format("2006-01-02")
		case "yesterday":
			dataDate["startDate"] = today.AddDate(0, 0, -1).Format("2006-01-02")
			// dataDate["endDate"] = today.AddDate(0, 0, -1).Format("2006-01-02")
		case "7":
			dataDate["startDate"] = today.AddDate(0, 0, -7).Format("2006-01-02")
			dataDate["endDate"] = today.Format("2006-01-02")
		case "this_month":
			dataDate["startDate"] = today.Format("2006-01")
		case "30":
			dataDate["startDate"] = today.AddDate(0, 0, -30).Format("2006-01-02")
			dataDate["endDate"] = today.Format("2006-01-02")
		}
	}

	conf := map[string]interface{}{
		"userId":     vars.Get("trans_user_id"),
		"date":       dataDate,
		"typeofDate": date,
	}

	dataResponse := controller.TransactionRepository.GetTransactionSummary(r.Context(), conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
