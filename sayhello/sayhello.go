package sayhello

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"httpHello/ports"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	html := `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
	<title>httpHello</title>
  </head>
  <body>
    <p>
       Hello!<br>
          %s

    </p>
  </body>
</html>
`
	url := r.URL.Path
	host := r.Host
	datetime := t.Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("\nTime is %s and you are calling from %s", datetime, host)
	// TODO: add 429 Too Many Requests & 418 I'm a teapot
	// https://en.wikipedia.org/wiki/Hyper_Text_Coffee_Pot_Control_Protocol
	switch url {
	case "/404":
		http.Error(w, "404 Not Found "+msg, http.StatusNotFound)
	case "/403":
		http.Error(w, "403 Forbidden "+msg, http.StatusForbidden)
	case "/408":
		http.Error(w, "408 Request Timeout "+msg, http.StatusRequestTimeout)
	case "/410":
		http.Error(w, "410 Gone "+msg, http.StatusGone)
	case "/418":
		http.Error(w, "418 I'm a teapot "+msg, http.StatusTeapot)
	case "/425":
		http.Error(w, "425 Too Early "+msg, http.StatusTooEarly)
	case "/429":
		http.Error(w, "429 Too Many Requests "+msg, http.StatusTooManyRequests)
	case "/500":
		http.Error(w, "500 Internal Server Error "+msg, http.StatusInternalServerError)
	case "/501":
		http.Error(w, "501 Not Implemented "+msg, http.StatusNotImplemented)
	case "/503":
		http.Error(w, "503 Service Unavailable "+msg, http.StatusServiceUnavailable)
	default:
		_, _ = fmt.Fprintf(w, html, msg)
	} // end switch

}

func StartListener(port string) error {

	// Is it an integer?
	n, err := strconv.Atoi(port)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "<%s> is not a valid port to listen on: %s", port, err)
		return err
	}

	// in valid port range?
	if n > 65535 || n < 1024 {
		_, _ = fmt.Fprintf(os.Stderr, "<%s> is not a valid portnumber", port)
		return err
	}

	if ports.Available(port) {
		http.HandleFunc("/", sayHello)

		err := http.ListenAndServe(":"+port, nil)

		if err != nil {
			return err
		}

	} else {
		e := fmt.Sprintf("Couldn't listen on port <%q>", port)
		_, _ = fmt.Fprintf(os.Stderr, e)
		return errors.New(e)
	}

	// no problems
	return nil
}
