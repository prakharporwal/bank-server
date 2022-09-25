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
import TransactionTable from "./Components/TransactionTable";
import { LoginForm, SignUpForm } from "./Components/LoginForm";
import StatementTable from "./Components/StatementTable";
import { AccountDetailsCard } from "./Components/AccountDetails";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Box textAlign="center" fontSize="xl">
      <Grid minH="100vh" p={3}>
        <ColorModeSwitcher justifySelf="flex-end" />
        <AccountDetailsCard />
        <VStack spacing={8}>
          <LoginForm></LoginForm>
          <Box height={220} width={1200} padding={4}>
            <TransactionTable />
            <StatementTable></StatementTable>
          </Box>
          {/* <SignUpForm></SignUpForm> */}
        </VStack>
      </Grid>
    </Box>
  </ChakraProvider>
);
