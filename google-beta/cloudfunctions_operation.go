package google

import (
	"fmt"
	"time"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/cloudfunctions/v1"
)

type CloudFunctionsOperationWaiter struct {
	Service *cloudfunctions.Service
	CommonOperationWaiter
}

func (w *CloudFunctionsOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	return w.Service.Operations.Get(w.Op.Name).Do()
}

func cloudFunctionsOperationWait(config *transport_tpg.Config, op *cloudfunctions.Operation, activity, userAgent string, timeout time.Duration) error {
	w := &CloudFunctionsOperationWaiter{
		Service: config.NewCloudFunctionsClient(userAgent),
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
