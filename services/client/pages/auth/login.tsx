import Link from "next/link";
import {useState} from "react";

export default () => {
    const [email, setEmail] = useState("");

    async function login () {
        const answer = await fetch("/api/internal/v1/auth/login", {
            method: "POST",
        })
    }

    return <>
        <h1>Login</h1>
        <input placeholder={"email"} value={email} onChange={(e) => setEmail(e.target.value)}/>
        <button onClick={login}>Click to login</button>
    </>
}