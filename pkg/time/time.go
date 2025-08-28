package time

import "time"

func FormatTimeToJakarta(t time.Time) time.Time {
    jakartaLocation, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        jakartaLocation = time.FixedZone("GMT+7", 7*60*60)
    }
    
    jakartaTime := t.In(jakartaLocation)
    return jakartaTime
}