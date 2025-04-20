export interface ReportModel {
  id: string
  latitude: number
  longitude: number
  location: string
  dateTime: Date
  timestamp: number
  timezone: number
  temperature: number
  pressure: number
  humidity: number
  cloudCover: number
  createdAt: Date
}

export interface ReportManyModel {
  list: ReportModel[]
}

export interface ReportCountManyModel {
  total: number
}
