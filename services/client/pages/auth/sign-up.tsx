import {useState} from "react";

export default () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [termsAndConditions, setTermsAndConditions] = useState(false);
    const [marketing, setMarketing] = useState(false);
    const [newsletter, setNewsletter] = useState(false);

    async function signup () {
        const answer = await fetch("/api/internal/v1/auth/register/magic-link", {
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
        const answer2 = await fetch("/api/internal/v1/auth/email-verification/send", {
            method: "POST",
            body: JSON.stringify({ "userId": answerJson.data.userId })
        })
    }

    return <>
        <h1>Sign up</h1>
        <input placeholder={"email"} value={email} onChange={(e) => setEmail(e.target.value)}/>
        <input placeholder={"username"} value={username} onChange={(e) => setUsername(e.target.value)}/>

        <br/>

        <input type={"checkbox"} checked={termsAndConditions} onClick={() => setTermsAndConditions((ps) => !ps)}/> Accept terms and conditions v1
        <br/>
        <input type={"checkbox"} checked={newsletter} onClick={() => setNewsletter((ps) => !ps)}/> Accept newsletter
        <br/>
        <input type={"checkbox"} checked={marketing} onClick={() => setMarketing((ps) => !ps)}/> Accept marketing

        <br/>

        <button onClick={signup}>Sign up</button>

        <br/>
        <a href={"/auth/email-verification/xxx"} target={"_blank"} />
    </>
}