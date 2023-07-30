import {useEffect} from "react";
import {useRouter} from "next/router";

export default (props) => {
    const router = useRouter()

    const { verificationCode } = router.query

    useEffect(() => {
        (async () => {
            if (verificationCode) {
                const unencodedVerificationCode = decodeUrlSafeBase64(verificationCode);
                const answer = await fetch("/api/v1/auth/email-verification/confirm", {
                    method: "PATCH",
                    body: JSON.stringify({ "verificationCode": unencodedVerificationCode })
                })
                const answerJson = await answer.json();
                if (answerJson.status !== "success") {
                    console.error(answerJson);
                    return;
                }
                const authorizationCode = answerJson.data.authorizationCode;
                const answer2 = await fetch("/api/v1/auth/token", {
                    method: "POST",
                    body: JSON.stringify({ "authorizationCode": authorizationCode })
                })
                const answer2Json = await answer2.json()
                if (answer2Json.status !== "success") {
                    console.error(answerJson);
                    return;
                }
                const accessToken = answer2Json.data.accessToken
                const refreshToken = answer2Json.data.refreshToken
                // store in local storage
                localStorage.setItem("accessToken", accessToken)
                localStorage.setItem("refreshToken", refreshToken)
            }
        })()
    }, [verificationCode])

    return <>...</>
}

function decodeUrlSafeBase64(encoded) {
    encoded = encoded.replaceAll(/-/g, '+').replaceAll(/_/g, '/');
    return Buffer.from(encoded, 'base64').toString();
}
