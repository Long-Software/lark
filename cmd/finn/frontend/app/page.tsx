import Image from "next/image";

export default function Home() {
  return (
    <div className="flex h-screen">
      {/* Main Content */}
      <main className="flex-1 p-6 grid grid-cols-3 grid-rows-2 gap-4">
        {/* Pie Chart of Expenses */}
        <div className="col-span-1 row-span-1 bg-white p-4 rounded-xl shadow">
          Pie Chart of Expense
        </div>

        {/* List of 5 Transactions */}
        <div className="col-span-1 row-span-1 bg-white p-4 rounded-xl shadow">
          List of 5 Transactions
        </div>

        {/* List of Category Expenses */}
        <div className="col-span-1 row-span-1 bg-white p-4 rounded-xl shadow">
          List of Categories Expense
        </div>

        {/* Area Chart of Income and Expense */}
        <div className="col-span-2 row-span-1 bg-white p-4 rounded-xl shadow">
          Area Chart of Income and Expense
        </div>

        {/* Extra Empty Section (for future use or customization) */}
        <div className="col-span-1 row-span-1 bg-white p-4 rounded-xl shadow">
          Extra Section
        </div>
      </main>
    </div>
  );
}
