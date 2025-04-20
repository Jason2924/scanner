import { CalendarIcon, TimerIcon } from '@radix-ui/react-icons'
import { Avatar, Box, Flex, Grid, Skeleton, Text } from '@radix-ui/themes'
import { ReportModel } from '../../models/report.model'
import { format } from 'date-fns'
import { getWeatherIcon } from '../../ultilities/common'

type Props = {
  data: ReportModel | undefined
}

const CurrentReport = ({data}: Props) => {
  return (
    <Box className="flex-2/3 text-white">
      <Box mb={"4"} py={"6"} px={"6"} className="bg-[#0005] rounded-lg">
        <Text as="p" mb={"2"} className="text-xl font-semibold">
          <Skeleton loading={data === undefined} width={"5rem"} height={"2rem"} />
          {data && data?.location}
        </Text>
        <Flex gap={"6"} align={"center"} justify={"center"}>
          <Skeleton loading={data === undefined} width={"15rem"} height={"6.5rem"} />
          {data && (
            <>
            <Box>
              <Text as="span" weight={"medium"} className="text-7xl">{Math.round(data.temperature)}</Text>
              <sup className="ml-2 text-5xl">°C</sup>
            </Box>
            <Box>
                <Avatar size={"5"} src={getWeatherIcon(data.dateTime)} fallback />
            </Box>
            </>
          )}
        </Flex>
        <span className="w-full h-[1px] mb-5 bg-gray-300 block"></span>
        <Box className="text-lg">
          <Flex gap={"3"} align={"center"} className="not-last:mb-2">
            <CalendarIcon className="inline-block size-6" />
            <Text as="span" className="inline-block text-gray-300 font-semibold">
              <Skeleton loading={data === undefined} width={"10rem"} height={"1.8rem"} />
              {data && format(data.dateTime, "MMM dd, yyyy")}
            </Text>
          </Flex>
          <Flex gap={"3"} align={"center"} className="not-last:mb-2">
            <TimerIcon className="inline-block size-6" />
            <Text as="span" className="inline-block text-gray-300 font-semibold">
              <Skeleton loading={data === undefined} width={"10rem"} height={"1.8rem"} />
              {data && format(data.dateTime, "hh:mm a - EE")}
            </Text>
          </Flex>
        </Box>
      </Box>
      <Grid gapX={"3"} columns={"3"}>
        <Box py={"4"} px={"5"} className="bg-[#0005] rounded-lg font-semibold">
          <Text as="p" mb={"1"} className="text-gray-300 text-lg">Pressure</Text>
          <Flex justify={"center"}>
            <Skeleton loading={data === undefined} width={"6rem"} height={"2rem"} />
          </Flex>
          {data && (
            <Text as="p" align={"center"} wrap={"nowrap"} className="text-2xl">
              {data.pressure}
              <span className="ml-2 text-gray-300 text-lg">hPa</span>
            </Text>
          )}
        </Box>
        <Box py={"4"} px={"5"} className="bg-[#0005] rounded-lg font-semibold">
          <Text as="p" mb={"1"} className="text-gray-300 text-lg">Humidity</Text>
          <Flex justify={"center"}>
            <Skeleton loading={data === undefined} width={"6rem"} height={"2rem"} />
          </Flex>
          {data && (
            <Text as="p" align={"center"} wrap={"nowrap"} className="text-2xl">
              {data.humidity}
              <span className="ml-2 text-gray-300 text-lg">%</span>
            </Text>
          )}
        </Box>
        <Box py={"4"} px={"5"} className="bg-[#0005] rounded-lg font-semibold">
          <Text as="p" mb={"1"} className="text-gray-300 text-lg">Cloud Cover</Text>
          <Flex justify={"center"}>
            <Skeleton loading={data === undefined} width={"6rem"} height={"2rem"} />
          </Flex>
          {data && (
            <Text as="p" align={"center"} wrap={"nowrap"} className="text-2xl">
              {data.cloudCover}
              <span className="ml-2 text-gray-300 text-lg">%</span>
            </Text>
          )}
        </Box>
      </Grid>
    </Box>
  )
}

export default CurrentReport
