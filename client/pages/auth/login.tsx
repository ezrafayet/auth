import Link from "next/link";

export default () => {

    async function login () {
        const answer = await fetch("/api/internal/v1/auth/login", {
            method: "POST",
        })
    }

    async function logout () {
        const answer = await fetch("/api/internal/v1/auth/logout", {
            method: "POST",
        })
    }

    return <>
        <h1>Login</h1>
        <button onClick={login}>Click to login</button>
        <button onClick={logout}>Click to logout</button>
        <Link href={"/auth/sign-up"}>
            <button>
                Sign up
            </button>
        </Link>
    </>
}