package tracer

import (
	"fmt"

	"go.elastic.co/apm/v2"
)

type ApmService struct {
	apmInstance *apm.Tracer
	Transaction *apm.Transaction
}

func (ap *ApmService) Initialize() {
	ap.apmInstance = apm.DefaultTracer()
}

func (ap *ApmService) StartTransaction() {

	ap.Transaction = ap.apmInstance.StartTransaction("test", "request")

	fmt.Println(ap.Transaction.TraceContext())

}

func (ap *ApmService) EndTransaction() {
	ap.Transaction.End()
}
