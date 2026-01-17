package constant

type ContactType string

const (
	Phone      ContactType = "phone"
	Whatsapp   ContactType = "whatsapp"
	Email      ContactType = "email"
	Telegram   ContactType = "telegram"
	OnlineShop ContactType = "online-shop"
	Instagram  ContactType = "instagram"
	Facebook   ContactType = "facebook"
	Tiktok     ContactType = "tiktok"
	Youtube    ContactType = "youtube"
)

type AddressType string

const (
	Home   AddressType = "home"
	Office AddressType = "office"
)
