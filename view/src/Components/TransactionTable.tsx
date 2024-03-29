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

type ITrans = {
  uid?: string;
  amount: number;
  created_at: string;
  from_account_id: number;
  to_account_id: number;
};

const TransactionTable: React.FunctionComponent<any> = () => {
  const [data, setData] = useState<ITrans[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/v1/transaction/list/1", {
      method: "GET",
      headers: new Headers({
        access_token: utils.getCookie("access_token"),
      }),
    })
      .then((response) => {
        console.log("res", response);
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
          <TableCaption>Transaction Tables</TableCaption>
          <Thead>
            <Tr>
              <Th>ID</Th>
              <Th>From Account Id</Th>
              <Th>To Account Id</Th>
              <Th>Amount</Th>
            </Tr>
          </Thead>
          <Tbody>
            {data.length !== 0 &&
              data.map((item) => {
                return (
                  <Tr key={item.uid}>
                    <Td>{item.uid?.split("-")[0]}</Td>
                    <Td>{item.from_account_id}</Td>
                    <Td>{item.to_account_id}</Td>
                    <Td isNumeric>{item.amount}</Td>
                  </Tr>
                );
              })}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  );
};

export default TransactionTable;
