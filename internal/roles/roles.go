package roles

type Role int64

const (
	Client Role = 1 << iota
	Seller
	Billing
	Manager
	Admin
)

func (r Role) IsAdmin() bool {
	return r&Admin != 0
}

func (r Role) IsManager() bool {
	return r&Manager != 0
}

func (r Role) IsBilling() bool {
	return r&Billing != 0
}

func (r Role) IsSeller() bool {
	return r&Seller != 0
}

func (r Role) IsClient() bool {
	return r&Client != 0
}

func (r Role) SetAdmin() Role {
	return r | Admin
}

func (r Role) SetManager() Role {
	return r | Manager
}

func (r Role) SetBilling() Role {
	return r | Billing
}

func (r Role) SetSeller() Role {
	return r | Seller
}

func (r Role) SetClient() Role {
	return r | Client
}

func (r Role) UnsetAdmin() Role {
	return r &^ Admin
}

func (r Role) UnsetManager() Role {
	return r &^ Manager
}

func (r Role) UnsetBilling() Role {
	return r &^ Billing
}

func (r Role) UnsetSeller() Role {
	return r &^ Seller
}

func (r Role) UnsetClient() Role {
	return r &^ Client
}

func (r Role) ToggleRole(role Role, enabled bool) Role {
	switch role {
	case Client:
		if enabled {
			return r.SetClient()
		} else {
			return r.UnsetClient()
		}
	case Manager:
		if enabled {
			return r.SetManager()
		} else {
			return r.UnsetManager()
		}
	case Billing:
		if enabled {
			return r.SetBilling()
		} else {
			return r.UnsetBilling()
		}
	case Seller:
		if enabled {
			return r.SetSeller()
		} else {
			return r.UnsetSeller()
		}
	case Admin:
		if enabled {
			return r.SetAdmin()
		} else {
			return r.UnsetAdmin()
		}
	}
	return r
}
