import TransactionListing from "@/components/TransactionListing";
import { ListTransaction } from "@/wailsjs/go/main/App";
import { transaction } from "@/wailsjs/go/models";
import { useEffect, useState } from "react";

const TransactionPage = () => {
  return (
    <div>
      <section>
        <TransactionListing />
      </section>
    </div>
  );
};

export default TransactionPage;
