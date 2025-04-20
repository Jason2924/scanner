import { Box, Container, Skeleton, Table, Text } from "@radix-ui/themes"
import axios from "axios";
import { format } from "date-fns"
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ReportManyModel, ReportModel } from "../models/report.model";

const ComparisonPage = () => {

  const REPORT_LENGTH = 2
  const { report1, report2 } = useParams<{ report1: string; report2: string }>();

  const [ reports, setReports ] = useState<ReportModel[]>([])
  const [ isLoading, setIsLoading ] = useState<boolean>(true)

  useEffect(() => {
    console.log(report1, report2)
      const fetchApi = async () => {
        try {
          await axios.get(
            "http://localhost:4000/api/v1/reports/compare-ids",
            {
              params: {ids: [report1, report2]}
            }
          ).then((result) => {
            const resultModel = result.data.data as ReportManyModel
            if (resultModel.list.length !== REPORT_LENGTH) {
              throw new Error(`Response length not equal to ${REPORT_LENGTH}`)
            }
            setReports(() => {
              const updated = resultModel.list.map((item) => ({...item, dateTime: new Date(item.dateTime)}))
              return [...updated]
            })
            setIsLoading(false)
          })
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
      fetchApi()
    }, [])

  return (
    <Container size={"3"} className="py-4 h-full">
      <Box py={"6"} px={"5"} className="text-white bg-[#0005] rounded-lg">
        <Text as="p" className="px-3 text-xl font-semibold">Comparison</Text>
        <Table.Root mb={"4"}>
          <Table.Header>
            <Table.Row>
              <Table.ColumnHeaderCell className="text-gray-300"></Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Report 1</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Report 2</Table.ColumnHeaderCell>
              <Table.ColumnHeaderCell className="text-gray-300">Deviation</Table.ColumnHeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            <Table.Row>
              <Table.RowHeaderCell className="text-gray-200">Timestamp</Table.RowHeaderCell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"6rem"} height={"3rem"} />
                {!isLoading && (
                  <>
                    <Text as="p" mb={"1"}>{format(reports[0].dateTime, "MMM dd, yyyy")}</Text>
                    <Text as="p">{format(reports[0].dateTime, "hh:mm a - EE")}</Text>
                  </>
                )}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"6rem"} height={"3rem"} />
                {!isLoading && (
                  <>
                    <Text as="p" mb={"1"}>{format(reports[1].dateTime, "MMM dd, yyyy")}</Text>
                    <Text as="p">{format(reports[1].dateTime, "hh:mm a - EE")}</Text>
                  </>
                )}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading}>-</Skeleton>
              </Table.Cell>
            </Table.Row>
            <Table.Row>
              <Table.RowHeaderCell className="text-gray-200">Temperature (°C)</Table.RowHeaderCell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.round(reports[0].temperature)}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.round(reports[1].temperature)}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.abs(Math.round(reports[0].temperature) - Math.round(reports[1].temperature))}
              </Table.Cell>
            </Table.Row>
            <Table.Row>
              <Table.RowHeaderCell className="text-gray-200">Pressure (hPa)</Table.RowHeaderCell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[0].pressure}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[1].pressure}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.abs(reports[0].pressure - reports[1].pressure)}
              </Table.Cell>
            </Table.Row>
            <Table.Row>
              <Table.RowHeaderCell className="text-gray-200">Humidity (%)</Table.RowHeaderCell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[0].humidity}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[1].humidity}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.abs(reports[0].humidity - reports[1].humidity)}
              </Table.Cell>
            </Table.Row>
            <Table.Row>
              <Table.RowHeaderCell className="text-gray-200">Cloud Cover (%)</Table.RowHeaderCell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[0].cloudCover}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && reports[1].cloudCover}
              </Table.Cell>
              <Table.Cell className="text-gray-200">
                <Skeleton loading={isLoading} width={"3rem"} height={"1.5rem"} />
                {!isLoading && Math.abs(reports[0].cloudCover - reports[1].cloudCover)}
              </Table.Cell>
            </Table.Row>
          </Table.Body>
        </Table.Root>
      </Box>
    </Container>
  )
}

export default ComparisonPage
