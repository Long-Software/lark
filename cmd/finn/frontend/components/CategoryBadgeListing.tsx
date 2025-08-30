"use client";
import { categories } from "@/data/categories";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import { Card } from "./ui/card";
import CreateCategoryDialog from "./CreateCategoryDialog";

const CategoryBadgeListing = () => {
  // const [categories, setCategories] = useState<category.Category[]>([]);
  // useEffect(() => {
  //   try {
  //     ListCategories().then((res) => setCategories(res));
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }, []);
  return (
    <div className="w-full">
      CategoryBadgeListing
      <div className="w-full flex justify-between px-10">
        <p className="text-xl">Categoires</p>
        {/* <Button variant="secondary">Add Category</Button> */}
        <CreateCategoryDialog />
      </div>
      <div className="w-full">
        <Card className="py-5 px-10">
          {categories.map((cat) => (
            <Badge key={cat.id} className="mx-1 my-0.5">
              {cat.name}
            </Badge>
          ))}
        </Card>
      </div>
    </div>
  );
};

export default CategoryBadgeListing;
