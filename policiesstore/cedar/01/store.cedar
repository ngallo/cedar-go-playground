permit (
    principal == User::"Alice",
    action == Action::"view",
    resource in Resource::"pharmacy-branch"
)
when { principal.age == 18 };

permit (
    principal == User::"Alice",
    action == Action::"view",
    resource in Album::"jane_vacation"
)
when { principal.age == 18 };