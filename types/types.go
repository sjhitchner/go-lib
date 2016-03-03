package types

import ()

// Date types
//go:generate jsondate -package types -name "DateYYMMDD" -format "060102"
//go:generate jsondate -package types -name "DateYYYYMMDD" -format "20060102"
//go:generate jsondate -package types -name "DateDYYYYMMDD" -format "2006-01-02"
//go:generate jsondate -package types -name "DateSYYMMDD" -format "06/01/02"
//go:generate jsondate -package types -name "DateSYYYYMMDD" -format "2006/01/02"
