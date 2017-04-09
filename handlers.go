package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/balaweblog/gosqlapi/model"
	"io/ioutil"
	"net/http"
)

var (
	// ErrInvalidInputRequest Error Invalid Request
	ErrInvalidInputRequest = errors.New("http: error reading input request")
	//ErrInvalidOutputResponse Error Invalid Output Response
	ErrInvalidOutputResponse = errors.New("http: error formatting response")
	//ErrParsingInputRequest Error Parsing Input Request
	ErrParsingInputRequest = errors.New("http: error parsing input request")
)

/*Createtable create table in sql */
func Createtable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "CREATE")
		val, err := model.ExecuteNonQuery(db, query)

		if val != nil && err == nil {
			output, _ := json.Marshal("table created successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Createdb create db in sql */
func Createdb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "CREATEDB")
		val, err := model.ExecuteNonQuery(db, query)

		if val != nil && err == nil {
			output, _ := json.Marshal("database created successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Showalltableindex create db in sql */
func Showalltableindex(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "SHOWALLTABLEINDEX")
		val, err := model.ExecuteNonQuery(db, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*Truncatetable  db in sql */
func Truncatetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "TRUNCATE")
		val, err := model.ExecuteNonQuery(db, query)

		if val != nil && err == nil {
			output, _ := json.Marshal("table truncated successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Droptable - drop table */
func Droptable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "DROPTABLE")
		val, err := model.ExecuteNonQuery(db, query)

		if val != nil && err == nil {
			output, _ := json.Marshal("table dropped successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Dropdatabase - drop database */
func Dropdatabase(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "DROPDB")
		val, err := model.ExecuteNonQuery(db, query)

		if val != nil && err == nil {
			output, _ := json.Marshal("database dropped successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Selecttable select from table in sql */
func Selecttable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "SELECT")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*Listallusers select from table in sql */
func Listallusers(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "LISTALLUSERS")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*Altertable alter table and columns in  sql */
func Altertable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "ALTER")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if val != nil && err == nil {
			output, _ := json.Marshal("table altered successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Inserttable insert into table in sql */
func Inserttable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "INSERT")
		val, err := model.ExecuteNonQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if val != nil && err == nil {
			output, _ := json.Marshal("records inserted successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Updatetable update into table in sql */
func Updatetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "UPDATE")
		val, err := model.ExecuteNonQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if val != nil && err == nil {
			output, _ := json.Marshal("update records successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Deletetable delete into table in sql */
func Deletetable(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "DELETE")
		val, err := model.ExecuteNonQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if val != nil && err == nil {
			output, _ := json.Marshal("deleted records successfully")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

/*Showalldatabases Show all databases in the sql */
func Showalldatabases(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "SHOWDB")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*CurrentInUseDB what database in use */
func CurrentInUseDB(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "CURRENTDB")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*ShowalltablesinDb show all tables in the database */
func ShowalltablesinDb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}

		query := model.ParseQuery(dat, "SHOWALLTABLES")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*Describetableindb describle individual table in the database */
func Describetableindb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "DESCTABLE")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

/*ExplaintableinDb explain table in sql */
func ExplaintableinDb(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || req == nil {
			http.Error(w, ErrInvalidInputRequest.Error(), http.StatusBadRequest)
			return
		}

		var dat map[string]interface{}
		err = json.Unmarshal(req, &dat)
		if err != nil {
			http.Error(w, ErrParsingInputRequest.Error(), http.StatusBadRequest)
			return
		}

		query := model.ParseQuery(dat, "EXPLAINTABLE")
		val, err := model.ExecuteQuery(db, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}
