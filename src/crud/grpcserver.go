package crud

import (
	"context"
	"fmt"
	"log"
	"simple-microservice/database"
	"strings"
)

// Релизация интерфейса gRPC сервера
type GRPCServer struct {}

func (s *GRPCServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	db := database.GetDB()

	where := ""
	if req.GetManufacturerId() > 0 {
		where = " and manufacturer_id = " + fmt.Sprint(req.GetManufacturerId())
	} else if req.GetId() > 0 {
		where = " and p.id = " + fmt.Sprint(req.GetId())
	}

	sql := "SELECT p.id, pm.name as manufacturer, p.vendor_code, p.created_at " +
		"FROM part p " +
		"LEFT JOIN part_manufacturer pm " +
		"ON pm.id = p.manufacturer_id " +
		"where p.deleted_at is null " +
		where
	rows, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var parts []*Part
	for rows.Next() {
		part := &Part{}
		err := rows.Scan(&part.Id, &part.Manufacturer, &part.VendorCode, &part.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		parts = append(parts, part)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &ListResponse{Parts: parts}, nil
}

func (s *GRPCServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	db := database.GetDB()

	values := ""
	var manIds []string

	for index, part := range req.GetParts() {
		if len(strings.Trim(part.GetVendorCode(), " ")) < 1 {
			return &CreateResponse{Status: CreateResponse_WRONG_VENDOR_CODE, Message: "Vendor code required"}, nil
		}

		values += fmt.Sprintf("(%d, '%s')", part.GetManufacturerId(), part.GetVendorCode())
		if index + 1 < len(req.GetParts()) {
			values += ","
		}
		manIds = append(manIds, fmt.Sprint(part.GetManufacturerId()))
	}

	var manCount int
	db.QueryRow("SELECT COUNT(*) as count from part_manufacturer where id in (" + strings.Join(manIds, ",") + ")").
		Scan(&manCount)

	if manCount != len(manIds) {
		return &CreateResponse{Status: CreateResponse_WRONG_MANUFACTURER, Message: "Manufacturer not found"}, nil
	}

	sqlInsert := "INSERT INTO part (\"manufacturer_id\", \"vendor_code\") VALUES " + values
	fmt.Println(sqlInsert)
	_, err := db.Exec(sqlInsert)

	if err != nil {
		return &CreateResponse{Status: CreateResponse_UNKNOWN_ERROR, Message: err.Error()}, nil
	}

	return &CreateResponse{Status: CreateResponse_SUCCESS, Message: "OK"}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	db := database.GetDB()
	for _, id := range(req.GetId()) {
		if id < 1 {
			return &DeleteResponse{Status: DeleteResponse_WRONG_PART_ID, Message: "Wrong part id: " + fmt.Sprint(id)}, nil
		}
	}

	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(req.GetId())), ","), "[]")
	_, err := db.Exec("update part set deleted_at = now() where id in (" + ids + ")")
	if err != nil {
		return &DeleteResponse{Status: DeleteResponse_OTHER_ERROR, Message: err.Error()}, nil
	}

	return &DeleteResponse{Status: DeleteResponse_SUCCESS, Message: "OK"}, nil
}