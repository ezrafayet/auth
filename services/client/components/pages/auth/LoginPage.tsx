import {useEffect, useState} from "react";
import Form from "../../form/Form";
import Spacer from "../../form/Spacer";
import TextInput from "../../form/TextInput";
import Checkbox from "../../form/Checkbox";
import Button from "../../form/Button";
import Link from "next/link";
import LinkButton from "../../form/LinkButton";
import {useRouter} from "next/router";

export default () => {
    const [email, setEmail] = useState("");

    const router = useRouter();

    async function login () {
        const answer = await fetch("/api/v1/auth/magic-link", {
            method: "POST",
            body: JSON.stringify({ email })
        })
        const answerJson = await answer.json();
        if (answerJson.status !== "success") {
            console.error(answerJson);
            return;
        }
    }

    return <div>

        <Spacer className={"mt-2 ml-3"}>
            <Link href={"/"}>
                <LinkButton text={"Back"} onClick={null} />
            </Link>
        </Spacer>

        <p>
            <br/>
        </p>

        <div className={"w-screen flex flex-col justify-center items-center"}>
        <Form onSubmit={(e) => {
            e.preventDefault()
            alert("Yo")
        }}>
            <h1 className={"my-5 font-bold text-lg"}>Login to your account</h1>

            <Spacer className={"flex justify-center items-center place-items-center text-sm w-full mt-3 mb-5"}>
                You will receive your email with a magic-link.
            </Spacer>

            <Spacer className={"my-2 w-full"}>
                <TextInput label={"email"}
                           placeHolder={""}
                           value={email}
                           type={"email"}
                           required={true}
                           onChange={(v: string) => setEmail(v)} fullWidth/>
            </Spacer>

            <Spacer className={"mb-2 mt-5 w-full"}>
                <Button text={"Login"} submit={true} fullWidth/>
            </Spacer>

            {/*<Spacer className={"text-pink-600 font-semibold text-sm"}>*/}
            {/*    <div>This is an error man</div>*/}
            {/*</Spacer>*/}

            <hr className={"border-gray-300 w-full mt-7 mb-3"}/>

            <Spacer className={"my-5 text-gray-400"}>
                Do not have an account? <Link href={`/auth/sign-up`}>
                <LinkButton text={"Sign up"}/>
            </Link>
            </Spacer>

        </Form>
    </div></div>
}