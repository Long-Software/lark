"use client";
import { category, transaction } from "@/wailsjs/go/models";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import CreateTransactionDialog from "./CreateTransactionDialog";
import { useEffect, useState } from "react";
import { ListCategories, ListTransaction } from "@/wailsjs/go/main/App";

const TransactionListing = ({}: {}) => {
  const [transactions, setTransactions] = useState<transaction.Transaction[]>(
    []
  );
  const [categories, setCategories] = useState<category.Category[]>([]);

  useEffect(() => {
    ListTransaction().then((res) => {
      setTransactions(res);
      console.log(res);
    });
    ListCategories().then(setCategories);
  }, []);
  const refresh = () => {
    ListTransaction().then(setTransactions);
  };
  return (
    <div className="w-full">
      CategoryBadgeListing
      <div className="w-full flex justify-between px-10">
        <p className="text-xl">Categoires</p>
        {/* <Button variant="secondary">Add Category</Button> */}
        <CreateTransactionDialog categories={categories} onConfirm={refresh} />
      </div>
      <div className="w-full">
        <Table>
          <TableCaption>A list of your recent invoices.</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px]">Invoice</TableHead>
              <TableHead className="text-center">Category</TableHead>
              <TableHead className="text-center">Amount</TableHead>
              <TableHead className="text-center">Timestamp</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {transactions.map((tran) => (
              <TableRow key={tran.id}>
                <TableCell className="font-medium">{tran.title}</TableCell>
                <TableCell className="text-center">
                  {tran.category.name}
                </TableCell>
                <TableCell className="text-center">{tran.amount}</TableCell>
                <TableCell className="text-center">{tran.created_at}</TableCell>
              </TableRow>
            ))}
          </TableBody>
          <TableFooter>
            <TableRow>
              <TableCell colSpan={3}>Total</TableCell>
              <TableCell className="text-right">$2,500.00</TableCell>
            </TableRow>
          </TableFooter>
        </Table>
      </div>
    </div>
  );
};

export default TransactionListing;
