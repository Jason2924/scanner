import { Box, Button, Container, Flex, Skeleton, Table, Text } from "@radix-ui/themes"
import { useEffect, useRef, useState } from "react"
import { ReportCountManyModel, ReportManyModel, ReportModel } from "../models/report.model"
import axios from "axios"
import { format } from "date-fns"
import { Checkbox } from "radix-ui"
import { CheckIcon } from "@radix-ui/react-icons"
import { useNavigate } from "react-router-dom"

const HistoryPage = () => {

  const API_URL = import.meta.env.VITE_API_URL

  const REPORT_LATITUDE = 1.3586
  const REPORT_LONGITUDE = 103.9899
  const REPORT_LIMIT = 5
  const MAX_SELECTION = 2

  const navigate = useNavigate();
  const firstApiRef = useRef(true);

  const [ histories, setHistories ] = useState<ReportModel[]>([])
  const [ current, setCurrent ] = useState<number>(1)
  const [ pages, setPages ] = useState<number>(1)
  const [ isLoading, setIsLoading ] = useState<boolean>(true)
  const [ selectedIds, setSelectedIds ] = useState<string[]>([])

  useEffect(() => {
    const fetchApi = async () => {
      try {
        await axios.get(
          `${API_URL}/reports/read-many`,
          {
            params: {
              latitude: REPORT_LATITUDE,
              longitude: REPORT_LONGITUDE,
              limit: REPORT_LIMIT,
              page: current
            }
          }
        ).then((result) => {
          const resultModel = result.data.data as ReportManyModel
          setHistories((previous) => {
            const updated = resultModel.list.map((item) => ({...item, dateTime: new Date(item.dateTime)}))
            if (previous) return [...previous!, ...updated]
            return [...updated]
          })
          setIsLoading(false)
        })
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    if (firstApiRef.current) {
      firstApiRef.current = false
      fetchApi()
    }
  }, [current])

  useEffect(() => {
    const fetchApi = async () => {
      try {
        await axios.get(
          `${API_URL}/reports/count-many`,
          {
            params: {
              latitude: REPORT_LATITUDE,
              longitude: REPORT_LONGITUDE
            }
          }
        ).then((result) => {
          const resultModel = result.data.data as ReportCountManyModel
          setPages(Math.ceil(resultModel.total/REPORT_LIMIT))
        })
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
    fetchApi()
  }, [])

  const handleShowMore = () => {
    setIsLoading(true)
    setCurrent(current+1)
    firstApiRef.current = true
  }

  const handleSelect = (checked: Checkbox.CheckedState, reportId: string) => {
    if (checked) {
      if (selectedIds.length < MAX_SELECTION) {
        setSelectedIds((previous) => [...previous, reportId])
        return
      }
      alert("Can not compare more than two reports")
      return
    }
    setSelectedIds((previous) => previous.filter((id) => id !== reportId))
  }

  const isChecked = (id: string) => selectedIds.includes(id);

  const handleCompare = () => {
    if (selectedIds.length === MAX_SELECTION) {
      const query = selectedIds.map((id) => id).join("/")
      navigate(`/comparison/${query}`)
    }
  } 

  return (
    <Container size={"3"} className="py-4 h-full">
      <Box py={"6"} px={"5"} className="text-white bg-[#0005] rounded-lg">
        <Flex justify={"between"} align={"center"}>
          <Text as="p" className="px-3 text-xl font-semibold">Histories</Text>
          <Button color="gray" variant="surface" highContrast onClick={handleCompare} disabled={selectedIds.length !== MAX_SELECTION} className="cursor-pointer">Compare</Button>
        </Flex>
        <Table.Root mb={"4"}>
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell className="text-gray-300">Date Time</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Temperature (°C)</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Pressure (hPa)</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Humidity (%)</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Cloud Cover (%)</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Select</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {!isLoading && histories ? (
              histories.map((history, index) => (
                <Table.Row key={index}>
                  <Table.RowHeaderCell className="text-gray-200">{format(history.dateTime, "MMM dd, yyyy | hh:mm a - EE")}</Table.RowHeaderCell>
                  <Table.Cell align={"center"} className="text-gray-200">{Math.round(history.temperature)}</Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">{history.pressure}</Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">{history.humidity}</Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">{history.cloudCover}</Table.Cell>
                  <Table.Cell align={"center"}>
                    <Checkbox.Root className="size-5 bg-white rounded flex items-center justify-center" checked={isChecked(history.id)} onCheckedChange={(checked) => handleSelect(checked, history.id)}>
                      <Checkbox.Indicator className="black">
                        <CheckIcon />
                      </Checkbox.Indicator>
                    </Checkbox.Root>
                  </Table.Cell>
                </Table.Row>
              ))
            ) : (
              Array.from({ length: 5 }).map((_, index) => (
                <Table.Row key={index}>
                  <Table.RowHeaderCell className="text-gray-200">
                    <Skeleton loading={isLoading} width={"12rem"} height={"1.5rem"} />
                  </Table.RowHeaderCell>
                  <Table.Cell align={"center"} className="text-gray-200">
                    <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                  </Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">
                    <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                  </Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">
                    <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                  </Table.Cell>
                  <Table.Cell align={"center"} className="text-gray-200">
                    <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                  </Table.Cell>
                  <Table.Cell align={"center"}>
                    <Skeleton loading={isLoading} width={"1.5rem"} height={"1.5rem"} />
                  </Table.Cell>
                </Table.Row>
              ))
            )}
          </Table.Body>
        </Table.Root>
        <Flex justify={"center"}>
          {current < pages && !isLoading && (
            <Button color="gray" variant="surface" onClick={handleShowMore} highContrast className="cursor-pointer">Show More</Button>
          )}
        </Flex>
      </Box>
    </Container>
  )
}

export default HistoryPage
