import nightIcon from "../assets/night.png"
import dayIcon from "../assets/day.png"

export const getWeatherIcon = (dateTime: Date) => {
  if (dateTime.getHours() > 5 && dateTime.getHours() < 18) {
    return dayIcon
  }
  return nightIcon
}
