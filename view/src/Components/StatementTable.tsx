import { useEffect, useState } from "react";
import {
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
} from "@chakra-ui/react";
import { utils } from "../utils/cookie";
import { error } from "console";

type IStatement = {
  account_id: number;
  amount: number;
  other_account: number;
  other_account_owner: string;
  timestamp: string;
  transaction_id: number;
  type: string;
};

const StatementTable: React.FunctionComponent<any> = () => {
  const [data, setData] = useState<IStatement[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/v1/account/3/statement/1", {
      method: "GET",
      headers: new Headers({
        access_token: utils.getCookie("access_token"),
      }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Something went wrong " + response.statusText);
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
        setData(data);
      })
      .catch((err) => {
        throw err;
      });
  }, []);

  return (
    <>
      <TableContainer>
        <Table variant="simple">
          <TableCaption>Account Statement</TableCaption>
          <Thead>
            <Tr>
              <Th>Transaction ID</Th>
              <Th>From Account </Th>
              <Th>Type</Th>
              <Th>Amount</Th>
              <Th>Timestamp</Th>
            </Tr>
          </Thead>
          <Tbody>
            {data.length !== 0 &&
              data.map((item) => {
                return (
                  <Tr key={item.transaction_id}>
                    <Td isNumeric>{item.transaction_id}</Td>
                    <Td>
                      {item.other_account + "\n" + item.other_account_owner}
                    </Td>
                    <Td>{item.type}</Td>
                    <Td isNumeric>{item.amount}</Td>
                    <Td>{item.timestamp}</Td>
                  </Tr>
                );
              })}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  );
};

export default StatementTable;
