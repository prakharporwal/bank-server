import { Box, Text, Skeleton, VStack } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { utils } from "../utils/cookie";

type Account = {
  id: string;
  owner_email: string;
  balance: number;
  currency: string;
};

const AccountDetailsCard: React.FunctionComponent<any> = (props) => {
  const [account, setAccount] = useState<Account>();
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(false);

  const ownerEmail = "prakharporwal@gmail.com";

  useEffect(() => {
    fetch("http://localhost:8080/account?owner_email=" + ownerEmail, {
      method: "GET",
      headers: new Headers({
        access_token: utils.getCookie("access_token"),
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setIsLoading(false);
        setAccount(data);
      })
      .catch((err) => {
        setError(true);
        throw err;
      });
  }, []);

  if (error) return <Box children="error" />;

  return (
    <Box height={24}>
      <VStack>
        <Skeleton height="20px" isLoaded={!isLoading}>
          <Text>{account?.owner_email}</Text>
        </Skeleton>
        <Skeleton height="20px" isLoaded={!isLoading}>
          <Text>{account?.balance}</Text>
        </Skeleton>
      </VStack>
    </Box>
  );
};

export { AccountDetailsCard };
