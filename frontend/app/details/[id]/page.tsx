import SinglePost from "@/components/SinglePost";

const Page = ({ params }: any) => {
  const { id } = params;
  console.log(id);
  return (
    <div>
      <SinglePost id={id} />
    </div>
  );
};

export default Page;
