permit (
    principal == User::"Alice",
    action == Action::"view",
    resource in Album::"adventure"
)
when { principal.age == 18 && 
    context.demoRequest == true };