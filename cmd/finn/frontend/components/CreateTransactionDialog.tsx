"use client";
import { toast } from "sonner";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { CreateTransaction } from "@/wailsjs/go/main/App";
import { FormEvent, useState } from "react";
import { DialogClose } from "@radix-ui/react-dialog";
import { category } from "@/wailsjs/go/models";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

const CreateTransactionDialog = ({
  categories,
  onConfirm,
}: {
  categories: category.Category[];
  onConfirm: () => void;
}) => {
  const [title, setTitle] = useState("");
  const [amount, setAmount] = useState(0);
  const [categoryID, setCategoryID] = useState(0);
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    try {
      CreateTransaction(title, amount, categoryID).catch((err) => {
        throw new Error(err);
      });
      toast(`${title}: ${amount} has been created`);
      onConfirm();
    } catch (error) {
      toast("Failed to create new category");
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">Create Transaction</Button>
      </DialogTrigger>

      <DialogContent className="sm:max-w-[425px]">
        <form onSubmit={handleSubmit}>
          <DialogHeader>
            <DialogTitle>Add New Trnasaction</DialogTitle>
            <DialogDescription>Add a new expense to your log</DialogDescription>
          </DialogHeader>

          <div className="grid gap-4 py-4">
            <div className="grid gap-3">
              <Label htmlFor="name">Title</Label>
              <Input
                id="name"
                name="title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
              />
            </div>
            <div className="grid gap-3">
              <Label htmlFor="amount">Amount</Label>
              <Input
                id="amount"
                type="number"
                name="title"
                value={amount}
                onChange={(e) => setAmount(Number(e.target.value))}
              />
            </div>
            <div className="grid gap-3">
              <Select onValueChange={(val) => setCategoryID(Number(val))}>
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Category" />
                </SelectTrigger>
                <SelectContent>
                  {categories.map((cat) => (
                    <SelectItem key={cat.id} value={cat.id.toString()}>
                      {cat.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <DialogFooter>
            <DialogClose asChild>
              <Button type="button" variant="outline">
                Cancel
              </Button>
            </DialogClose>
            <Button type="submit">Confirm</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateTransactionDialog;
