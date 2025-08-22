package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SubscriptionCreateParameters represents the SubscriptionCreateParameters schema from the OpenAPI specification
type SubscriptionCreateParameters struct {
	Secondarykey string `json:"secondaryKey,omitempty"` // Secondary subscription key. If not specified during request key will be generated automatically.
	State string `json:"state,omitempty"` // Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	Userid string `json:"userId"` // User (user id path) for whom subscription is being created in form /users/{uid}
	Name string `json:"name"` // Subscription name.
	Primarykey string `json:"primaryKey,omitempty"` // Primary subscription key. If not specified during request key will be generated automatically.
	Productid string `json:"productId"` // Product (product id path) for which subscription is being created in form /products/{productId}
}

// SubscriptionUpdateParameters represents the SubscriptionUpdateParameters schema from the OpenAPI specification
type SubscriptionUpdateParameters struct {
	State string `json:"state,omitempty"` // Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	Statecomment string `json:"stateComment,omitempty"` // Comments describing subscription state change by the administrator.
	Userid string `json:"userId,omitempty"` // User identifier path: /users/{uid}
	Expirationdate string `json:"expirationDate,omitempty"` // New subscription expiration date.
	Name string `json:"name,omitempty"` // Subscription name.
	Primarykey string `json:"primaryKey,omitempty"` // Primary subscription key.
	Productid string `json:"productId,omitempty"` // Product identifier path: /products/{productId}
	Secondarykey string `json:"secondaryKey,omitempty"` // Secondary subscription key.
}

// SubscriptionCollection represents the SubscriptionCollection schema from the OpenAPI specification
type SubscriptionCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []SubscriptionContract `json:"value,omitempty"` // Page values.
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
}

// SubscriptionContract represents the SubscriptionContract schema from the OpenAPI specification
type SubscriptionContract struct {
	Createddate string `json:"createdDate,omitempty"` // Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Enddate string `json:"endDate,omitempty"` // Date when subscription was cancelled or expired. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Productid string `json:"productId,omitempty"` // The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
	Secondarykey string `json:"secondaryKey,omitempty"` // Subscription secondary key.
	Startdate string `json:"startDate,omitempty"` // Subscription activation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expirationdate string `json:"expirationDate,omitempty"` // Subscription expiration date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Id string `json:"id,omitempty"` // Uniquely identifies the subscription within the current API Management service instance. The value is a valid relative URL in the format of /subscriptions/{sid} where {sid} is a subscription identifier.
	Notificationdate string `json:"notificationDate,omitempty"` // Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Name string `json:"name,omitempty"` // The name of the subscription, or null if the subscription has no name.
	Statecomment string `json:"stateComment,omitempty"` // Optional subscription comment added by an administrator.
	Userid string `json:"userId,omitempty"` // The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
	Primarykey string `json:"primaryKey,omitempty"` // Subscription primary key.
	State string `json:"state,omitempty"` // Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
}
