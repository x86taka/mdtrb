package bgpview

type BGPViewIP struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          Data   `json:"data"`
	Meta          Meta   `json:"@meta"`
}
type Asn struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}
type Prefixes struct {
	Prefix      string `json:"prefix"`
	IP          string `json:"ip"`
	Cidr        int    `json:"cidr"`
	Asn         Asn    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}
type RirAllocation struct {
	RirName          string `json:"rir_name"`
	CountryCode      string `json:"country_code"`
	IP               string `json:"ip"`
	Cidr             string `json:"cidr"`
	Prefix           string `json:"prefix"`
	DateAllocated    string `json:"date_allocated"`
	AllocationStatus string `json:"allocation_status"`
}
type IanaAssignment struct {
	AssignmentStatus string      `json:"assignment_status"`
	Description      string      `json:"description"`
	WhoisServer      string      `json:"whois_server"`
	DateAssigned     interface{} `json:"date_assigned"`
}
type Maxmind struct {
	CountryCode string      `json:"country_code"`
	City        interface{} `json:"city"`
}
type Data struct {
	IP             string         `json:"ip"`
	PtrRecord      string         `json:"ptr_record"`
	Prefixes       []Prefixes     `json:"prefixes"`
	RirAllocation  RirAllocation  `json:"rir_allocation"`
	IanaAssignment IanaAssignment `json:"iana_assignment"`
	Maxmind        Maxmind        `json:"maxmind"`
}
type Meta struct {
	TimeZone      string `json:"time_zone"`
	APIVersion    int    `json:"api_version"`
	ExecutionTime string `json:"execution_time"`
}

type GETASN struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Asn               int      `json:"asn"`
		Name              string   `json:"name"`
		DescriptionShort  string   `json:"description_short"`
		DescriptionFull   []string `json:"description_full"`
		CountryCode       string   `json:"country_code"`
		Website           string   `json:"website"`
		EmailContacts     []string `json:"email_contacts"`
		AbuseContacts     []string `json:"abuse_contacts"`
		LookingGlass      string   `json:"looking_glass"`
		TrafficEstimation string   `json:"traffic_estimation"`
		TrafficRatio      string   `json:"traffic_ratio"`
		OwnerAddress      []string `json:"owner_address"`
		RirAllocation     struct {
			RirName          string `json:"rir_name"`
			CountryCode      string `json:"country_code"`
			DateAllocated    string `json:"date_allocated"`
			AllocationStatus string `json:"allocation_status"`
		} `json:"rir_allocation"`
		IanaAssignment struct {
			AssignmentStatus string      `json:"assignment_status"`
			Description      string      `json:"description"`
			WhoisServer      string      `json:"whois_server"`
			DateAssigned     interface{} `json:"date_assigned"`
		} `json:"iana_assignment"`
		DateUpdated string `json:"date_updated"`
	} `json:"data"`
	Meta struct {
		TimeZone      string `json:"time_zone"`
		APIVersion    int    `json:"api_version"`
		ExecutionTime string `json:"execution_time"`
	} `json:"@meta"`
}
