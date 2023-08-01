
export default function Spacer (props: {
    className: string
    children?: any
}) {

    return <div className={props.className}>
        {props.children}
    </div>
}