import Link from "next/link";
import {useState} from "react";

export default () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");

    async function signup () {
        const answer = await fetch("/api/internal/v1/auth/register", {
            method: "POST",
            body: JSON.stringify({ method: "magicLink", email, username })
        })
    }

    return <>
        <h1>Sign up</h1>
        <input placeholder={"email"} value={email} onChange={(e) => setEmail(e.target.value)}/>
        <input placeholder={"username"} value={username} onChange={(e) => setUsername(e.target.value)}/>

        <button onClick={signup}>Sign up</button>
    </>
}