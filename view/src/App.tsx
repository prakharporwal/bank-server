import * as React from "react";
import {
  ChakraProvider,
  Box,
  Text,
  Link,
  VStack,
  Code,
  Grid,
  theme,
} from "@chakra-ui/react";
import { ColorModeSwitcher } from "./ColorModeSwitcher";
import { Logo } from "./Logo";
import TransactionTable from "./Components/TransactionTable";
import LoginForm from "./Components/LoginForm";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Box textAlign="center" fontSize="xl">
      <Grid minH="100vh" p={3}>
        <ColorModeSwitcher justifySelf="flex-end" />
        <VStack spacing={8}>
          <Box height={220} width={600} padding={4}>
            <LoginForm></LoginForm>
          </Box>
          <TransactionTable />
        </VStack>
      </Grid>
    </Box>
  </ChakraProvider>
);
