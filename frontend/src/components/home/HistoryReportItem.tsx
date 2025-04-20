import { Avatar, Flex, Text } from '@radix-ui/themes'
import { ReportModel } from '../../models/report.model'
import { format } from 'date-fns'
import { getWeatherIcon } from '../../ultilities/common'

type Props = {
  data: ReportModel
  handleClick: (id: string) => void
}

const HistoryReportItem = ({data, handleClick}: Props) => {
  return (
    <Flex px={"3"} py={"2"} align={"center"} justify={"between"} onClick={() => handleClick(data.id)} className="cursor-pointer hover:bg-[#0007] bg-[#0005] rounded-xl text-lg text-gray-300">
      <Flex gap={"3"} align={"center"}>
        <Avatar src={getWeatherIcon(data.dateTime)} fallback size={"2"} />
        <Text>{format(data.dateTime, "hh:mm a")}</Text>
      </Flex>
      <Text className="text-xl font-semibold">{Math.round(data.temperature)} <sup className="text-sm">°C</sup></Text>
    </Flex>
  )
}

export default HistoryReportItem
