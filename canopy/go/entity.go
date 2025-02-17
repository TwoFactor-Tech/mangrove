package mangrove

type Result struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type Model struct {
	ID        string `json:"id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	Version   int64  `json:"version,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type NamedModel struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type OwnedModel struct {
	User  User `json:"user,omitempty"`
	Owner User `json:"Owner,omitempty"`
}

type UserCredential struct {
	Model
	NamedModel
	OwnedModel

	ExternalId string `json:"externalId,omitempty"`
	Secret     string `json:"secret,omitempty"`
}

type MerchantOrder struct {
	Model
	NamedModel
	OwnedModel

	Merchant             User   `json:"merchant,omitempty"`
	Order                Order  `json:"order,omitempty"`
	ExternalIPForCreator string `json:"externalIPForCreator,omitempty"`
	ExternalIPForPayer   string `json:"externalIPForPayer,omitempty"`
}

type Order struct {
	Model
	NamedModel
	OwnedModel

	OrderTransactionId string        `json:"orderTransactionId,omitempty"`
	Status             string        `json:"status,omitempty"`
	Amount             string        `json:"amount,omitempty"`
	IntegralAmount     string        `json:"integralAmount,omitempty"`
	ExternalId         string        `json:"externalId,omitempty"`
	Code               string        `json:"code,omitempty"`
	OrderCallback      OrderCallback `json:"orderCallback,omitempty"`
}

type OrderCallback struct {
	Model

	Endpoint string `json:"endpoint,omitempty"`
}

type User struct {
	Model
	NamedModel

	Type       string `json:"type,omitempty"`
	LogStatus  string `json:"logStatus,omitempty"`
	Integral   string `json:"integral,omitempty"`
	LastSeenAt string `json:"lastSeenAt,omitempty"`
}

type UserOrder struct {
	Model
	NamedModel
	OwnedModel

	Currency string `json:"currency,omitempty"`
	Order    Order  `json:"order,omitempty"`
}

type UserWallet struct {
	Model
	NamedModel
	OwnedModel

	Status               string                `json:"status,omitempty"`
	Type                 string                `json:"type,omitempty"`
	AcceptableRange      []any                 `json:"acceptableRange,omitempty"`
	LastUsedAt           string                `json:"lastUsedAt,omitempty"`
	UserWalletAttributes []UserWalletAttribute `json:"userWalletAttributes,omitempty"`
}

type UserWalletAttribute struct {
	Model
	OwnedModel

	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
