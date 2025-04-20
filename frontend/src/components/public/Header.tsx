import { Box, Container, Flex } from "@radix-ui/themes"
import { Link } from "react-router-dom"

const Header = () => {
  return (
    <header className="py-3 shadow">
      <Container size={"3"}>
        <Flex justify={"between"}>
          <Box className="flex-1/12 max-w-32">
            <Link className="block" to="/">Scanner</Link>
          </Box>
          <Box className="flex-auto">
            <Flex gap={"3"} align={"center"} justify={"end"}>
              <Link className="px-2 block" to="/">Home</Link>
              <Link className="px-2 block" to="/history">History</Link>
            </Flex>
          </Box>
        </Flex>
      </Container>
    </header>
  )
}

export default Header
