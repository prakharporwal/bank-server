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

type ITrans = {
  id?: string;
  amount: number;
  created_at: string;
  from_account_id: number;
  to_account_id: number;
};

const TransactionTable: React.FunctionComponent<any> = () => {
  const [data, setData] = useState<ITrans[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/v1/transaction/list/1")
      .then((response) => {
        // console.log("res", response);
        return response.json();
      })
      .then((data) => {
        console.log(data.slice(0, 1));
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
            {data.map((item) => {
              return (
                <Tr key={item.id}>
                  <Td>{item.id}</Td>
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
