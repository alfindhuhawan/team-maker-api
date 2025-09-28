# Directory
- folder name /airpax-consult
- you must download gogen3.exe to run gogen command
- put your gogen3.exe outside your directory /airpax-consult, that make you call the gogen3.exe on cmd with this command "./gogen3.exe"

# Gogen Command
- ../gogen3.exe init myprofile<Nama Domain> //nama domain tidak boleh ada spasi
- ../gogen3.exe usecase RunAdminCreate<Nama Usecase> //nama usecase tidak ada spasi
    > Pada inport.go edit inport requestnya berdasarkan data yang diperlukan
    > Pada interactor.go tulis logicnya
    > Pada outport.go tulis keperluan model model yang dipanggil
->>> kalo usecase buat POST / PUT depannya harus Run ex : ../gogen3.exe usecase RunAdminCreate
->>> kalo usecase buat GET depannya harus Get ex : ../gogen3.exe usecase GetAdminCreate
- ../gogen3.exe repository SaveAdmin<Nama fungsi di repository> Admin<Nama Entity> RunAdminCreate<Nama usecase>
    > tulis nama tabel di entity
- ../gogen3.exe error <Nama validasi> //nama tidak ada spasi, camelcase
- ../gogen3.exe gateway prod<Nama gateway>
    > tulis cara koneksi ke db
- ../gogen3.exe controller restapi
    > edit untuk mengatur input / output data
- go mod tidy
- ../gogen3.exe application myperson
- go run main.go myperson

# Rule Directory
- All function to call from database or external source, put it on shared/gateway, with "impl_" name folder
- After creating impl_ set repository or service to call the function on interactor
- use repository if the impl_ function get data from database (/shared/model/repository)
- use service if the impl_ function get data from external source (/shared/model/service)