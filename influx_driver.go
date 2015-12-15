package main 

import( 
		// "github.com/influxdb/influxdb/client/v2"
)

func QueryInfluxDB(clnt client.Client, cmd string) (res []client.Result, err error) {


    q := client.Query{
        Command:  cmd,
        Database: db,
    }


    if response, err := clnt.Query(q); err == nil {
        if response.Error() != nil {
            return res, response.Error()
        }
        res = response.Results
    }


    return res, nil
} 

func WriteInfluxDB(c client.Client, string data ) (err error) {
	var err error

	return err
}
