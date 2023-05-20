import Link from "next/link";
import {useState} from "react";

export default () => {
    const [email, setEmail] = useState("");

    async function login () {
        const answer = await fetch("/api/internal/v1/auth/magic-link", {
            method: "POST",
            body: JSON.stringify({ email })
        })
        const answerJson = await answer.json();
        if (answerJson.status !== "success") {
            console.error(answerJson);
            return;
        }
    }

    return <>
        <h1>Login</h1>
        <input placeholder={"email"} value={email} onChange={(e) => setEmail(e.target.value)}/>
        <button onClick={login}>Click to login</button>
        <br/>
        <a href={"/auth/login/xxx"} target={"_blank"} />
    </>
}