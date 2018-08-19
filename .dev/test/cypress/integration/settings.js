context('Settings Forms', () => {

    beforeEach(function() {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should edit main settings', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get('input[name="project"]').clear().type('Project Updated')
        cy.get('input[name="description"]').clear().type('This is an awesome page')
        cy.get('input[name="domain"]').clear().type('http://0.0.0.0:8080')
        cy.get('textarea[name="footer"]').clear().type('This is a custom footer')
        cy.get('#v-pills-home > form').submit()
        cy.title().should('eq', 'Statup | Settings')
        cy.get('input[name="project"]').should('have.value', 'Project Updated')
        cy.get('input[name="description"]').should('have.value', 'This is an awesome page')
        cy.get('input[name="domain"]').should('have.value', 'http://0.0.0.0:8080')
        cy.get('.footer').should('contain', 'This is a custom footer')
    })

    // it('should check index page for changes', () => {
    //     cy.visit('http://localhost:8080/')
    //     cy.title().should('eq', 'Project Updated Status')
    //     cy.get('.header-title').should('contain', 'Project Updated')
    //     cy.get('.header-desc').should('contain', 'This is an awesome page')
    // })

});