import Link from "next/link";

export default () => {
    async function logout () {
        const answer = await fetch("/api/internal/v1/auth/logout", {
            method: "POST",
        })
    }

    return <>
        <h1>Home</h1>
        <Link href={"/auth/sign-up"}>
            <button>
                Sign Up
            </button>
        </Link>
        <Link href={"/auth/login"}>
            <button>
                Login
            </button>
        </Link>
        <button onClick={logout}>
            Logout
        </button>
    </>
}
