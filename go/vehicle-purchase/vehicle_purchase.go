
package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
  if kind == "truck" || kind == "car" {
    return true
  } else {
    return false
  }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    text := " is clearly the better choice."
	if option1 < option2 {
        return option1 + text
    } else {
        return option2 + text
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
  if int(age) >= 3 && int(age) < 10 {
      return originalPrice * 0.70
  } else if int(age) >= 10 {
      return originalPrice * 0.50
  }

  return originalPrice * 0.80
}
