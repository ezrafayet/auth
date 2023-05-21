import {useState} from "react";

export default () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");

    async function signup () {
        const answer = await fetch("/api/internal/v1/auth/register", {
            method: "POST",
            body: JSON.stringify({ "authType": "magic-link", email, username })
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

        <button onClick={signup}>Sign up</button>

        <br/>
        <a href={"/auth/email-verification/xxx"} target={"_blank"} />
    </>
}