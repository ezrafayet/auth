import {useRouter} from "next/router";
import {useEffect} from "react";

export default (props) => {
    const router = useRouter()

    const { authorizationCode } = router.query

    useEffect(() => {
        (async () => {
            if (authorizationCode) {
                const unencodedCode = decodeUrlSafeBase64(authorizationCode);
                const answer = await fetch("/api/v1/auth/token", {
                    method: "POST",
                    body: JSON.stringify({ "authorizationCode": unencodedCode })
                })
                const answerJson = await answer.json();
                if (answerJson.status !== "success") {
                    console.error(answerJson);
                    return;
                }
                const accessToken = answerJson.data.accessToken
                const refreshToken = answerJson.data.refreshToken
                // store in local storage
                localStorage.setItem("accessToken", accessToken)
                localStorage.setItem("refreshToken", refreshToken)
            }
        })()
    }, [authorizationCode])

    return <>...</>
}

function decodeUrlSafeBase64(encoded) {
    encoded = encoded.replaceAll(/-/g, '+').replaceAll(/_/g, '/');
    return Buffer.from(encoded, 'base64').toString();
}
