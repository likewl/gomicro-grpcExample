cd services/protos
protoc  --proto_path=. --micro_out=../ --go_out=../ models.proto
protoc  --proto_path=. --micro_out=../ --go_out=../ Prodservices.proto
cd ..&cd..