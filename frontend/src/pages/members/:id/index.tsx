import { format, formatRFC3339, parseISO } from "date-fns";
import { useCallback, useEffect, useState, VFC } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import { MembersClassForm } from "../../../components/forms/MembersClassForm";
import { MembersClassList } from "../../../components/MembersClassList";
import {
  Gender,
  UpdateMemberInput,
  useGetMemberLazyQuery,
  useUpdateMemberMutation,
} from "../../../generated/graphql";

type Inputs = {
  id: number;
  number: number;
  name: string;
  kana: string;
  gender: Gender;
  dateOfBirth: string;
  phoneNumber: string;
  email: string;
  dateOfAdmission: string;
  dateOfWithdrawal: string;
  memo: string;
};

export const MemberPage: VFC = () => {
  const { memberId } = useParams();
  const { register, handleSubmit, setValue } = useForm<Inputs>();
  const [getMember, { data, loading, error }] = useGetMemberLazyQuery();
  const [updateMember] = useUpdateMemberMutation();
  const [editable, setEditable] = useState<boolean>(false);
  const [showMembersClassForm, setShowMembersClassForm] =
    useState<boolean>(false);

  useEffect(() => {
    if (!memberId) return;

    getMember({ variables: { id: memberId } });
  }, [memberId, getMember]);

  const setValueFromdata = useCallback(() => {
    if (!data) return;
    const { member } = data;
    setValue("id", parseInt(member.id));
    setValue("number", member.number);
    setValue("name", member.name);
    setValue("kana", member.kana);
    setValue("gender", member.gender);
    setValue("dateOfBirth", format(parseISO(member.dateOfBirth), "yyyy-MM-dd"));
    setValue("phoneNumber", member.phoneNumber || "");
    setValue("email", member.email || "");
    setValue(
      "dateOfAdmission",
      format(parseISO(member.dateOfAdmission), "yyyy-MM-dd")
    );
    setValue(
      "dateOfWithdrawal",
      member.dateOfWithdrawal
        ? format(parseISO(member.dateOfWithdrawal), "yyyy-MM-dd")
        : ""
    );
    setValue("memo", member.memo);
    setGender(member.gender);
  }, [data, setValue]);

  useEffect(() => {
    setValueFromdata();
  }, [setValueFromdata]);

  const [gender, setGender] = useState<Gender | null>(null);

  const renderCancelButton = () => {
    if (editable) {
      return (
        <button
          type="button"
          onClick={() => {
            setEditable(!editable);
            setValueFromdata();
          }}
        >
          ???????????????
        </button>
      );
    } else {
      return <div></div>;
    }
  };

  const renderEditButton = () => {
    if (editable) {
      return (
        <button type="submit" onClick={handleSubmit(onSubmit)}>
          ??????
        </button>
      );
    } else {
      return (
        <button type="button" onClick={() => setEditable(true)}>
          ??????
        </button>
      );
    }
  };

  const renderButtons = () => {
    return (
      <div className="flex flex-row justify-between">
        {renderCancelButton()}
        {renderEditButton()}
      </div>
    );
  };

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    const input: UpdateMemberInput = {
      ...data,
      id: data.id.toString(),
      dateOfBirth: data.dateOfBirth
        ? formatRFC3339(parseISO(data.dateOfBirth))
        : null,
      dateOfAdmission: formatRFC3339(parseISO(data.dateOfAdmission)),
      dateOfWithdrawal: data.dateOfWithdrawal
        ? formatRFC3339(parseISO(data.dateOfWithdrawal))
        : null,
    };
    updateMember({ variables: { input } })
      .then((member) => getMember({ variables: { id: data.id.toString() } }))
      .finally(() => setEditable(false));
  };

  if (!data) return <></>;
  const { member } = data;
  if (!member) return <></>;

  return (
    <>
      <h1>??????</h1>
      <form className="flex flex-col">
        <label htmlFor="number">??????</label>
        <input
          {...register("number", { required: true, disabled: !editable })}
        />
        <label htmlFor="name">??????</label>
        <input {...register("name", { required: true, disabled: !editable })} />
        <label htmlFor="kana">????????????</label>
        <input {...register("kana", { required: true, disabled: !editable })} />
        <label htmlFor="gender">??????</label>
        <div className="flex flex-row items-center">
          ???
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Female.toString()}
            onChange={() => setGender(Gender.Female)}
            checked={gender === Gender.Female}
          />
          ???
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Male.toString()}
            onChange={() => setGender(Gender.Male)}
            checked={gender === Gender.Male}
          />
          ?????????
          <input
            {...register("gender", {
              required: true,
              disabled: !editable,
            })}
            type="radio"
            value={Gender.Other.toString()}
            onChange={() => setGender(Gender.Other)}
            checked={gender === Gender.Other}
          />
        </div>
        <label htmlFor="dateOfBirth">????????????</label>
        <input
          type="date"
          {...register("dateOfBirth", { disabled: !editable })}
        />
        <label htmlFor="phoneNumber">????????????</label>
        <input
          type="tel"
          {...register("phoneNumber", { disabled: !editable })}
        />
        <label htmlFor="email">?????????</label>
        <input type="email" {...register("email", { disabled: !editable })} />
        <label htmlFor="dateOfAdmission">?????????</label>
        <input
          type="date"
          {...register("dateOfAdmission", { disabled: !editable })}
        />
        <label htmlFor="dateOfWithdrawal">?????????</label>
        <input
          type="date"
          {...register("dateOfWithdrawal", { disabled: !editable })}
        />
        <label htmlFor="memo">??????</label>
        <input type="text" {...register("memo", { disabled: !editable })} />
      </form>
      {renderButtons()}
      <hr />
      <h1>?????????</h1>
      {showMembersClassForm ? (
        <button onClick={() => setShowMembersClassForm(false)}>
          ???????????????
        </button>
      ) : (
        <button onClick={() => setShowMembersClassForm(true)}>??????</button>
      )}
      <div className={showMembersClassForm ? "" : "hidden"}>
        <MembersClassForm
          member={member}
          onSubmitted={() => setShowMembersClassForm(false)}
        />
      </div>
      <div className={showMembersClassForm ? "hidden" : ""}>
        <MembersClassList membersClasses={member.membersClasses} />
      </div>
    </>
  );
};
