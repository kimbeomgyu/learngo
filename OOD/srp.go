package OOD

type Report struct {
}

type FinanceReport struct {
}

func (r *FinanceReport) MakeReport() *Report {
	// ...
}

type ReportSender interface {
	SendReport(*Report)
}

type EmailReportSender struct {
}

func (r *EmailReportSender) SendReport(r *Report) {
	// .. email
}

type FileReportSender struct {
}

func (s *FileReportSender) SendReport(r *Report) {
	// .. File
}
