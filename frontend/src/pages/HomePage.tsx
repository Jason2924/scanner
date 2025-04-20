import { Box, Container, Flex, Grid, Skeleton, Text } from "@radix-ui/themes"
import CurrentReport from "../components/home/CurrentReport"
import { ReportManyModel, ReportModel } from "../models/report.model"
import { useEffect, useState } from "react"
import axios from "axios"
import HistoryReportItem from "../components/home/HistoryReportItem"

const HomePage = () => {

  const API_URL = import.meta.env.VITE_API_URL

  const REPORT_LATITUDE = 1.3586
  const REPORT_LONGITUDE = 103.9899
  const REPORT_LIMIT = 5
  const REPORT_PAGE = 1

  const [ current, setCurrent ] = useState<ReportModel>()
  const [ histories, setHistories ] = useState<ReportModel[]>()

  useEffect(() => {
    const fetchApi = async () => {
      try {
        const currentReportApi = axios.get(
          `${API_URL}/reports/read-current`,
          {
            params: {
              latitude: 1.3586,
              longitude: 103.9899
            }
          }
        ).then((result) => result.data.data as ReportModel)
        const historyReportApi = axios.get(
          `${API_URL}/reports/read-many`,
          {
            params: {
              latitude: REPORT_LATITUDE,
              longitude: REPORT_LONGITUDE,
              limit: REPORT_LIMIT,
              page: REPORT_PAGE
            }
          }
        ).then((result) => result.data.data as ReportManyModel)
        const [ currentResult, historyResult ] = await Promise.all([currentReportApi, historyReportApi])
        setCurrent({...currentResult, dateTime: new Date(currentResult.dateTime)})
        setHistories(historyResult.list.map((history) => ({...history, dateTime: new Date(history.dateTime)})))
      } catch (error) {
        console.error('Error fetching data:', error);
      } 
    }
    fetchApi()
  }, [])

  const handleClick = (id: string) => {
    if (id !== current?.id) {
      const newCurrent = histories?.find((history) => history.id === id)
      setCurrent(newCurrent)
    }
  }

  return (
    <Container size={"3"} className="py-4 h-full">
      <Flex gap={"3"}>
        <CurrentReport data={current} />
        <Box py={"6"} px={"5"} className="flex-1/3 text-white bg-[#0005] rounded-lg">
          <Text as="p" mb={"4"} className="px-3 text-xl font-semibold">Histories</Text>
          <Grid gapY={"3"}>
            {histories ? (
              histories.map((history, index) => (
                <HistoryReportItem key={index} data={history} handleClick={handleClick} />
              ))
            ) : (
              Array.from({ length: 5 }).map((_, index) => (
                <Skeleton key={index} loading={true} width={"15rem"} height={"3rem"} />
              ))
            )}
          </Grid>
        </Box>
      </Flex>
    </Container>
  )
}

export default HomePage
