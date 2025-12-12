export class ErrorState {
    error: boolean = $state(false)
    message: string = $state("")

    Error() {
        this.error = true;
    }

    Success() {
        this.error = false;
    }

    SetMessage(msg: string) {
        this.message = msg;
    }
}
