import { TemplateResult, html } from "lit";
import { customElement, property } from "lit/decorators.js";

import { EVENT_REFRESH } from "../../common/constants";
import "../LoadingOverlay";
import { ModalButton } from "../buttons/ModalButton";
import "../buttons/SpinnerButton";
import { Form } from "../forms/Form";

@customElement("ak-forms-modal")
export class ModalForm extends ModalButton {
    @property({ type: Boolean })
    closeAfterSuccessfulSubmit = true;

    @property({ type: Boolean })
    showSubmitButton = true;

    @property()
    submitKeepOpen?: string;

    @property({ type: Boolean })
    loading = false;

    @property({ type: String })
    cancelText = "Cancel";

    async confirm(): Promise<void> {
        const form = this.querySelector<Form<unknown>>("[slot=form]");
        if (!form) {
            return Promise.reject("No form found");
        }
        const formPromise = form.submit(new Event("submit"));
        if (!formPromise) {
            return Promise.reject("Form didn't return a promise for submitting");
        }
        return formPromise
            .then(() => {
                if (this.closeAfterSuccessfulSubmit) {
                    this.open = false;
                    form?.resetForm();
                }
                this.loading = false;
                this.locked = false;
                this.dispatchEvent(
                    new CustomEvent(EVENT_REFRESH, {
                        bubbles: true,
                        composed: true,
                    }),
                );
            })
            .catch((exc) => {
                this.loading = false;
                this.locked = false;
                throw exc;
            });
    }

    renderModalInner(): TemplateResult {
        return html`${this.loading
                ? html`<ak-loading-overlay ?topMost=${true}></ak-loading-overlay>`
                : html``}
            <section class="pf-c-modal-box__header pf-c-page__main-section pf-m-light">
                <div class="pf-c-content">
                    <h1 class="pf-c-title pf-m-2xl">
                        <slot name="header"></slot>
                    </h1>
                </div>
            </section>
            <section class="pf-c-modal-box__body pf-c-page__main-section pf-m-light">
                <slot name="form"></slot>
            </section>
            <footer class="pf-c-modal-box__footer">
                ${this.showSubmitButton
                    ? html`<ak-spinner-button
                              .callAction=${() => {
                                  this.loading = true;
                                  this.locked = true;
                                  return this.confirm();
                              }}
                              class="pf-m-primary"
                          >
                              <slot name="submit"></slot> </ak-spinner-button
                          >&nbsp;`
                    : html``}
                ${this.submitKeepOpen
                    ? html`
                          <ak-spinner-button
                              .callAction=${() => {
                                  this.loading = true;
                                  this.locked = true;
                                  this.closeAfterSuccessfulSubmit = false;
                                  return this.confirm().finally(() => {
                                      this.closeAfterSuccessfulSubmit = true;
                                  });
                              }}
                              class="pf-m-primary"
                          >
                              <slot name="${this.submitKeepOpen}"></slot> </ak-spinner-button
                          >&nbsp;
                      `
                    : html``}
                <ak-spinner-button
                    .callAction=${async () => {
                        this.resetForms();
                        this.open = false;
                    }}
                    class="pf-m-secondary"
                >
                    ${this.cancelText}
                </ak-spinner-button>
            </footer>`;
    }
}
