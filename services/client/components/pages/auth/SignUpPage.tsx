import {useEffect, useState} from "react";
import Form from "../../form/Form";
import TextInput from "../../form/TextInput";
import Button from "../../form/Button";
import Link from "next/link";
import LinkButton from "../../form/LinkButton";
import Checkbox from "../../form/Checkbox";
import Spacer from "../../form/Spacer";
import {useRouter} from "next/router";

export default () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [termsAndConditions, setTermsAndConditions] = useState(false);
    const [marketing, setMarketing] = useState(false);
    const [improveService, setImproveService] = useState(false);
    const [newsletter, setNewsletter] = useState(false);

    const router = useRouter();

    useEffect(() => {
        const email = router.query.email;
        setEmail(email as string);
    }, [])

    const [success, setSuccess] = useState(false);
    const [code, setCode] = useState("");

    async function signup() {
        const answer = await fetch("/api/v1/auth/register", {
            method: "POST",
            body: JSON.stringify({
                "authType": "magic-link",
                email,
                username,
                hasAcceptedTerms: termsAndConditions,
                acceptedTermsVersion: "v1",
                hasAcceptedMarketing: marketing,
                hasAcceptedNewsletter: newsletter
            })
        })
        const answerJson = await answer.json();
        if (answerJson.status !== "success") {
            console.error(answerJson);
            return;
        }
        await fetch("/api/v1/auth/email-verification/send", {
            method: "POST",
            body: JSON.stringify({"userId": answerJson.data.userId})
        })
        setSuccess(true);
    }

    if (success) {
        return <div>
            <h1>Enter the code you got in an email:</h1>
            <input placeholder={"code"} value={code} onChange={(e) => setCode(e.target.value)}/>
            <br/>
            <button onClick={async () => {
            }}>Verify
            </button>
        </div>
    }

    return <div className={"w-screen flex flex-col justify-center items-center"}>

        <p>
            <br/>
        </p>

        <Form onSubmit={(e) => {
            e.preventDefault()
            alert("Yo")
        }}>
            <h1 className={"my-5 font-bold text-lg"}>Create an account</h1>

            <Spacer className={"flex justify-center items-center place-items-center text-sm w-full mt-3 mb-5"}>
                Note a valid email is needed to verify your account.
            </Spacer>


            <Spacer className={"my-2 w-full"}>
                <TextInput label={"email"}
                           placeHolder={""}
                           value={email}
                           type={"email"}
                           required={true}
                           onChange={(v: string) => setEmail(v)} fullWidth/>
            </Spacer>

            <Spacer className={"my-2 w-full"}>
                <TextInput label={"username"}
                           placeHolder={""}
                           value={username}
                           required={true}
                           onChange={(v: string) => setUsername(v)} fullWidth/>
            </Spacer>

            <Spacer className={"mb-1 mt-4 w-full"}>
                <Spacer className={"my-2 w-full"}>
                    <Checkbox text={"I wish to receive marketing emails"}
                              checked={marketing}
                              onChange={() => setMarketing((ps) => !ps)}/>
                </Spacer>

                <Spacer className={"my-2 w-full"}>
                    <Checkbox text={"I have read and approved terms and conditions and privacy policy"}
                              checked={termsAndConditions}
                              required={true}
                              onChange={() => setTermsAndConditions((ps) => !ps)}/>
                </Spacer>
            </Spacer>

            <Spacer className={"mb-2 mt-5 w-full"}>
                <Button text={"Sign Up"} submit={true} fullWidth/>
            </Spacer>

            <hr className={"border-gray-300 w-full mt-7 mb-3"}/>

            <Spacer className={"my-5 text-gray-400"}>
                Already have an account? <Link href={`/auth/login?email=${email}`}>
                <LinkButton text={"Login"}/>
            </Link>
            </Spacer>

        </Form>

        <br/><br/>
    </div>
}