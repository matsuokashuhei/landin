import { useEffect, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import { Layout } from "../../../components";
import {
  DeleteInstructorInput,
  Instructor,
  UpdateInstructorInput,
  useDeleteInstructorMutation,
  useGetInstructorLazyQuery,
  useUpdateInstructorMutation,
} from "../../../generated/graphql";

type Inputs = {
  id: number;
  name: string;
  syllabicCharacters: string;
  biography: string;
  phoneNumber: string;
  email: string;
};

export const InstructorPage: VFC = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [editable, setEditable] = useState<Boolean>(false);
  const [getInstructor, { data, loading, error }] = useGetInstructorLazyQuery();
  const [updateInstructor] = useUpdateInstructorMutation();
  const [deleteInstructor] = useDeleteInstructorMutation();

  const { register, handleSubmit } = useForm<Inputs>();

  useEffect(() => {
    if (id) {
      getInstructor({ variables: { id } });
    }
  }, [id, getInstructor]);

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateInstructorInput = {
      id: data.id.toString(),
      name: data.name,
      syllabicCharacters: data.syllabicCharacters,
      biography: data.biography,
      phoneNumber: data.phoneNumber,
      email: data.email,
    };
    updateInstructor({ variables: { input } })
      .then(() => getInstructor({ variables: { id: data.id.toString() } }))
      .then(() => setEditable(false));
  };

  const renderInstructor = (instructor: Instructor) => {
    if (editable) {
      return (
        <form className="flex flex-col">
          <input
            type="hidden"
            {...register("id", { required: true })}
            defaultValue={instructor.id}
          />
          <label htmlFor="name">名前</label>
          <input
            {...register("name", { required: true })}
            defaultValue={instructor.name}
          />
          <label htmlFor="syllabicCharacters">よみがな</label>
          <input
            {...register("syllabicCharacters", { required: true })}
            defaultValue={instructor.syllabicCharacters}
          />
          <label htmlFor="biography">紹介文</label>
          <input
            {...register("biography")}
            defaultValue={instructor.biography ?? ""}
          />
          <label htmlFor="phoneNumber">電話番号</label>
          <input
            {...register("phoneNumber")}
            defaultValue={instructor.phoneNumber ?? ""}
          />
          <label htmlFor="email">メール</label>
          <input {...register("email")} defaultValue={instructor.email ?? ""} />
        </form>
      );
    } else {
      return (
        <div className="flex flex-col">
          <div>{instructor.id}</div>
          <div>{instructor.name}</div>
          <div>{instructor.syllabicCharacters}</div>
          <div>{instructor.biography}</div>
          <div>{instructor.phoneNumber}</div>
          <div>{instructor.email}</div>
          <div>{instructor.createTime}</div>
          <div>{instructor.updateTime}</div>
        </div>
      );
    }
  };

  const renderEditButton = () => {
    if (editable) {
      return (
        <>
          <button type="button" onClick={() => setEditable(false)}>
            取消
          </button>
          <button type="submit" onClick={handleSubmit(onSubmit)}>
            保存
          </button>
        </>
      );
    } else {
      return <button onClick={() => setEditable(true)}>編集</button>;
    }
  };

  const renderDeleteButton = () => {
    if (editable) return <></>;
    if (!data) return <></>;
    const { id }: Instructor = data?.instructor;

    return (
      <button
        onClick={() => {
          const input: DeleteInstructorInput = {
            id: id,
          };
          deleteInstructor({ variables: { input } }).then(() =>
            navigate("/instructors")
          );
        }}
      >
        削除
      </button>
    );
  };

  if (!data) return <></>;
  const { instructor } = data;

  return (
    <Layout>
      {renderInstructor(instructor)}
      {renderEditButton()}
      {renderDeleteButton()}
    </Layout>
  );
};