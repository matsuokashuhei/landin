import { format, formatRFC3339, parseISO } from "date-fns";
import { useEffect, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import {
  GetClassQuery,
  UpdateClassInput,
  useDeleteClassMutation,
  useGetClassLazyQuery,
  useUpdateClassMutation,
} from "../../../generated/graphql";

type Clazz = GetClassQuery["class"];

type Inputs = {
  id: number;
  name: string;
  level: string;
  tuition: number;
  startDate: string;
  endDate: string;
  instructorId: number;
  scheduleId: number;
};

export const ClassPage: VFC = () => {
  const { classId } = useParams();
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<Inputs>();
  const [getClass, { data, loading, error }] = useGetClassLazyQuery();
  const [UpdateClass] = useUpdateClassMutation();
  const [DeleteClass] = useDeleteClassMutation();
  const [clazz, setClazz] = useState<Clazz>();
  const [editable, setEditable] = useState<boolean>(false);

  useEffect(() => {
    if (classId) {
      getClass({ variables: { id: classId } });
    }
  }, [classId, getClass]);

  useEffect(() => {
    if (data) {
      setClazz(data.class);
    }
  }, [data]);

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateClassInput = {
      id: data.id.toString(),
      name: data.name,
      level: data.level,
      tuition: data.tuition,
      startDate: formatRFC3339(parseISO(data.startDate)),
      endDate: data.endDate ? formatRFC3339(parseISO(data.endDate)) : null,
      instructorId: data.instructorId.toString(),
      scheduleId: data.scheduleId.toString(),
    };
    UpdateClass({ variables: { input } })
      .then((data) => navigate(`/classes/`))
      .catch((error) => console.error(error));
  };

  const handleDelete = (id: string) => {
    DeleteClass({ variables: { id: id } })
      .then((data) => navigate(`/classes/`))
      .catch((error) => console.error(error));
  };

  if (!clazz) return <></>;

  console.log(clazz);

  return (
    <>
      <h1>?????????</h1>
      <form className="flex flex-col">
        <input
          type="hidden"
          {...register("id", { required: true })}
          defaultValue={clazz.id}
        />
        <label htmlFor="name">??????</label>
        <input
          type="text"
          defaultValue={clazz.name}
          disabled={!editable}
          {...register("name", { required: true })}
        />
        <label htmlFor="level">?????????</label>
        <input
          type="text"
          defaultValue={clazz.level}
          disabled={!editable}
          {...register("level", { required: true })}
        />
        <label htmlFor="tuition">??????</label>
        <input
          type="number"
          defaultValue={clazz.tuition}
          disabled={!editable}
          {...register("tuition", { required: true })}
        />
        <label htmlFor="startDate">?????????</label>
        <input
          type="date"
          defaultValue={format(parseISO(clazz.startDate), "yyyy-MM-dd")}
          disabled={!editable}
          {...register("startDate", { required: true })}
        />
        <label htmlFor="endDate">?????????</label>
        <input
          type="date"
          defaultValue={
            clazz.endDate && format(parseISO(clazz.endDate), "yyyy-MM-dd")
          }
          disabled={!editable}
          {...register("endDate", { required: false })}
        />
        <label htmlFor="instructorId">????????????????????????</label>
        <select {...register("instructorId")}>
          <option value={clazz.instructor.id} disabled={!editable} selected>
            {clazz.instructor.name}
          </option>
        </select>
        <input
          type="hidden"
          {...register("scheduleId", { required: true })}
          defaultValue={clazz.schedule.id}
        />
      </form>
      {editable ? (
        <div className="flex flex-row justify-between">
          <button type="button" onClick={() => setEditable(!editable)}>
            ???????????????
          </button>
          <button type="submit" onClick={handleSubmit(onSubmit)}>
            ??????
          </button>
        </div>
      ) : (
        <div className="flex flex-row justify-between">
          <button type="button" onClick={() => handleDelete(clazz.id)}>
            ??????
          </button>
          <button type="button" onClick={() => setEditable(!editable)}>
            ??????
          </button>
        </div>
      )}
    </>
  );
};
