import Link from "next/link";

export default () => {
    return <>
        <h1>Sign up</h1>
        <Link href={"/auth/login"}>
            <button>
                Login
            </button>
        </Link>
    </>
}