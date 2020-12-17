package mysql

// Notes
/* https://github.com/arnehormann/sqlinternals
   https://godoc.org/github.com/arnehormann/sqlinternals/mysqlinternals
    https://github.com/go-sql-driver/mysql/issues/585
    https://go-review.googlesource.com/c/go/+/29961

*/

var columnPointers []interface{}
